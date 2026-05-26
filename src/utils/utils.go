package utils

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-redis/redis/v8"
	"github.com/go-sql-driver/mysql"

	. "Notifier/models"
)

var ErrorLogger *log.Logger
var SentNoticeLogger *log.Logger
var PostLogger *log.Logger
var DB *sql.DB
var ctx = context.Background()

// Redis에서 crawling-token 가져오기
func GetTokenFromRedis() string {
	// Redis 클라이언트 설정
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")

	rdb := redis.NewClient(&redis.Options{
		Addr: redisHost + ":" + redisPort, // 환경변수에서 Redis 호스트 정보 가져오기
	})

	// crawling-token 키로 Redis에서 토큰 가져오기
	token, err := rdb.Get(ctx, "crawling-token").Result()
	if err != nil {
		log.Fatalf("Failed to get token from Redis: %v", err)
	}
	return token
}

func CreateLogger(writer io.Writer) *log.Logger {
	return log.New(writer, "", log.Ldate|log.Ltime|log.Lshortfile)
}

func LogInfo(event string, fields ...any) {
	writeLog(SentNoticeLogger, "info", event, fields...)
}

func LogError(event string, fields ...any) {
	writeLog(ErrorLogger, "error", event, fields...)
}

func LogPost(event string, fields ...any) {
	writeLog(PostLogger, "info", event, fields...)
}

func writeLog(logger *log.Logger, level string, event string, fields ...any) {
	if logger == nil {
		logger = log.Default()
	}

	var builder strings.Builder
	builder.WriteString("level=")
	builder.WriteString(level)
	builder.WriteString(" event=")
	builder.WriteString(event)

	for i := 0; i+1 < len(fields); i += 2 {
		key, ok := fields[i].(string)
		if !ok || key == "" {
			continue
		}
		builder.WriteByte(' ')
		builder.WriteString(key)
		builder.WriteByte('=')
		builder.WriteString(formatLogValue(fields[i+1]))
	}

	logger.Println(builder.String())
}

func formatLogValue(value any) string {
	switch v := value.(type) {
	case string:
		return fmt.Sprintf("%q", v)
	case error:
		return fmt.Sprintf("%q", v.Error())
	case time.Duration:
		return fmt.Sprintf("%d", v.Milliseconds())
	default:
		return fmt.Sprint(v)
	}
}

func ConnectDB() *sql.DB {
	startedAt := time.Now()
	config := mysql.Config{
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PW"),
		Net:                  "tcp",
		Addr:                 os.Getenv("DB_IP") + ":" + os.Getenv("DB_PORT"),
		DBName:               os.Getenv("DB_NAME"),
		AllowNativePasswords: true,
	}
	connector, err := mysql.NewConnector(&config)
	if err != nil {
		ErrorLogger.Panic(err)
	}
	db := sql.OpenDB(connector)
	err = db.Ping()
	if err != nil {
		LogError("db_connect_failed", "error", err, "elapsed_ms", time.Since(startedAt))
		ErrorLogger.Panic(err)
	}
	LogInfo("db_connected", "elapsed_ms", time.Since(startedAt))
	return db
}

func LoadNotifierConfig(path string) []NotifierConfig {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var configs []NotifierConfig
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configs)
	if err != nil {
		log.Fatal(err)
	}
	return configs
}

func LoadDbData(topic string) (int, int) {
	boxCount := loadNoticeValue(topic, "box")
	maxNum := loadNoticeValue(topic, "num")

	return boxCount, maxNum
}

func loadNoticeValue(topic, noticeType string) int {
	var value int
	query := "SELECT n.value FROM notice AS n JOIN topics AS t ON n.topic_id = t.id WHERE t.department = ? AND n.type = ?"

	err := DB.QueryRow(query, topic, noticeType).Scan(&value)
	if err == nil {
		return value
	}
	if err != sql.ErrNoRows {
		log.Fatal(err)
	}

	insertQuery := "INSERT INTO notice (topic_id, type, value) SELECT t.id, ?, ? FROM topics AS t WHERE t.department = ?"
	result, err := DB.Exec(insertQuery, noticeType, 0, topic)
	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	if rowsAffected == 0 {
		log.Fatalf("topic %q does not exist in topics table", topic)
	}

	return 0
}

// 웹훅을 호출할 때 Redis에서 가져온 토큰을 Bearer로 헤더에 추가
func SendCrawlingWebhook(url string, payload any) {
	startedAt := time.Now()
	payloadJson, err := json.Marshal(payload)
	if err != nil {
		ErrorLogger.Panic(err)
	}
	buff := bytes.NewBuffer(payloadJson)

	// Redis에서 crawling-token 가져오기
	token := GetTokenFromRedis()

	// HTTP 요청 생성
	req, err := http.NewRequest("POST", url, buff)
	if err != nil {
		ErrorLogger.Panic(err)
	}

	// Content-Type 헤더 설정
	req.Header.Set("Content-Type", "application/json")

	// Authorization 헤더에 Bearer 토큰 설정
	req.Header.Set("crawling-token", token)

	// HTTP 클라이언트로 요청 보내기
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		LogError("webhook_finished", "result", "failed", "error", err, "elapsed_ms", time.Since(startedAt))
		ErrorLogger.Panic(err)
	}
	defer resp.Body.Close()

	// 응답 본문 읽기
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ErrorLogger.Panic(err)
	}
	result := "success"
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		result = "failed"
	}
	LogPost("webhook_finished", "result", result, "status", resp.StatusCode, "elapsed_ms", time.Since(startedAt), "response_bytes", len(body))
}

func GetNumNoticeCountReference(doc *goquery.Document, englishTopic, boxNoticeSelector string) int {
	if englishTopic != "Software" {
		return 10
	}
	boxNoticeSels := doc.Find(boxNoticeSelector)
	boxCount := boxNoticeSels.Length()
	return 15 - boxCount
}

func NewDocumentFromPage(url string) (*goquery.Document, error) {
	startedAt := time.Now()
	// HTTP GET 요청을 위한 새로운 요청 생성
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// 요청 생성 에러 발생 시 에러 반환
		LogError("http_request_create_failed", "error", err)
		return nil, err
	}

	// User-Agent 헤더 설정
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36")

	// HTTP 클라이언트 생성
	client := &http.Client{}

	// 요청 실행
	resp, err := client.Do(req)
	if err != nil {
		// 네트워크 에러 발생 시 에러 반환
		LogError("http_request_failed", "result", "failed", "error", err, "elapsed_ms", time.Since(startedAt))
		return nil, err
	}

	defer resp.Body.Close()

	// 응답 상태 코드 확인
	if resp.StatusCode != http.StatusOK {
		// 상태 코드 에러 발생 시 에러 반환
		LogError("http_request_finished", "result", "failed", "status", resp.StatusCode, "elapsed_ms", time.Since(startedAt))
		return nil, fmt.Errorf("status code error: %d, URL: %s", resp.StatusCode, url)
	}

	// HTML 문서 파싱
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		// 파싱 에러 발생 시 에러 반환
		LogError("html_parse_failed", "result", "failed", "error", err, "elapsed_ms", time.Since(startedAt))
		return nil, err
	}

	return doc, nil
}
