# AjouEvent Crawling

아주대학교 공지사항을 주기적으로 수집해 AjouEvent 서비스로 전달하는 크롤링 서비스입니다.

## 크롤링 대상

이 서비스는 아주대학교 대표 공지, 장학 공지, 기숙사 공지와 각 단과대학 및 학과 홈페이지의 공지사항 게시판을 대상으로 합니다.

### 학교 공통

- 아주대학교 일반 공지
- 아주대학교 장학 공지
- 기숙사 공지
- 대학원 공지

### 단과대학 및 학부

- 경영대학
- 공과대학
- 국제학부대학
- 다산학부대학
- 사회과학대학
- 소프트웨어융합대학
- 약학대학
- 인문대학
- 자연과학대학
- 첨단ICT융합대학
- 첨단바이오융합대학
- 간호대학
- 의과대학
- 프런티어과학학부
- 자유전공학부
- 경제정치사회융합학부

### 학과 및 전공

- AI모빌리티공학과
- 건설시스템공학과
- 건축학과
- 경영인텔리전스학과
- 경영학과
- 경제학과
- 교통시스템공학과
- 국방디지털융합학과
- 국어국문학과
- 글로벌경영학과
- 금융공학과
- 기계공학과
- 디지털미디어학과
- 문화콘텐츠학과
- 물리학과
- 불어불문학과
- 사학과
- 사회학과
- 산업공학과
- 생명과학과
- 소프트웨어학과
- 수학과
- 스포츠레저학과
- 심리학과
- 영어영문학과
- 응용화학과
- 응용화학생명공학과
- 인공지능융합학과
- 전자공학과
- 정치외교학과
- 지능형반도체공학과
- 첨단신소재공학과
- 사이버보안학과
- 융합시스템공학과
- 행정학과
- 화학공학과
- 화학과
- 환경안전공학과

## 대상 URL

| 대상 | URL |
|---|---|
| AI모빌리티공학과 | https://mobility.ajou.ac.kr/mobility/board/notice.do |
| 간호대학 | https://www.ajoumc.or.kr/nursing/board/commBoardNRNewsList.do |
| 건설시스템공학과 | https://ce.ajou.ac.kr/ce/board/notice.do |
| 건축학과 | https://arch.ajou.ac.kr/arch/board/total-notice.do |
| 경영대학 | https://biz.ajou.ac.kr/biz/community/notice.do |
| 경영인텔리전스학과 | https://ebiz.ajou.ac.kr/ebiz/board/notice.do |
| 경영학과 | https://abiz.ajou.ac.kr/abiz/board/notice.do |
| 경제정치사회융합학부 | https://eps.ajou.ac.kr/eps/life/notice.do |
| 경제학과 | https://econ.ajou.ac.kr/econ/activity/notice.do |
| 공과대학 | https://eng.ajou.ac.kr/eng/community/notice.do |
| 교통시스템공학과 | https://tse.ajou.ac.kr/tse/board/board03.do |
| 국방디지털융합학과 | https://mdc.ajou.ac.kr/mdc/administration/notice.do |
| 국어국문학과 | https://kor.ajou.ac.kr/kor/academic/notice.do |
| 국제학부대학 | https://isa.ajou.ac.kr/isa/support/notice.do |
| 글로벌경영학과 | https://gb.ajou.ac.kr/gb/admission/alram.do |
| 금융공학과 | https://fe.ajou.ac.kr/fe/notice/notice.do |
| 기계공학과 | https://me.ajou.ac.kr/me/board/under-notice.do |
| 기숙사 | https://dorm.ajou.ac.kr/dorm/notice/notice.do |
| 다산학부대학 | https://uc.ajou.ac.kr/uc/comunity/cnotice.do |
| 대학원 | https://grad.ajou.ac.kr/gs/community/notice.do |
| 디지털미디어학과 | https://media.ajou.ac.kr/media/board/notice.do |
| 문화콘텐츠학과 | https://culture.ajou.ac.kr/culture/academic/notice.do |
| 물리학과 | https://physics.ajou.ac.kr/physics/board/notice.do |
| 불어불문학과 | https://france.ajou.ac.kr/france/notice/notice.do |
| 사이버보안학과 | https://security.ajou.ac.kr/security/board/under-notice.do |
| 사학과 | https://history.ajou.ac.kr/history/admission/notice.do |
| 사회과학대학 | https://coss.ajou.ac.kr/coss/community/community01.do |
| 사회학과 | https://soci.ajou.ac.kr/soci/activity/notice.do |
| 산업공학과 | https://ie.ajou.ac.kr/ie/academic/notice.do |
| 생명과학과 | https://biology.ajou.ac.kr/biolog/admission/notice.do |
| 소프트웨어융합대학 | https://sw.ajou.ac.kr/sw/board/notice.do |
| 소프트웨어학과 | http://software.ajou.ac.kr/bbs/board.php?tbl=notice |
| 수학과 | https://math.ajou.ac.kr/math/alram/notice.do |
| 스포츠레저학과 | https://slez.ajou.ac.kr/slez/activity/notice.do |
| 심리학과 | https://apsy.ajou.ac.kr/apsy/activity/notice.do |
| 아주대학교 일반 공지 | https://ajou.ac.kr/kr/ajou/notice.do |
| 아주대학교 장학 공지 | https://ajou.ac.kr/kr/ajou/notice_scholarship.do |
| 약학대학 | https://pharm.ajou.ac.kr/pharm/board/notice.do |
| 영어영문학과 | https://ell.ajou.ac.kr/english/dgree/notice.do |
| 융합시스템공학과 | https://ise.ajou.ac.kr/ise/board/notice.do |
| 응용화학과 | https://appchem.ajou.ac.kr/appchem/college/notice.do |
| 응용화학생명공학과 | https://chembio.ajou.ac.kr/chembio/admission/notice.do |
| 의과대학 | https://www.ajoumc.or.kr/medicine/board/commBoardUVNoticeList.do |
| 인공지능융합학과 | https://aai.ajou.ac.kr/aai/dgree/notice.do |
| 인문대학 | https://human.ajou.ac.kr/human/community/community01.do |
| 자연과학대학 | https://ns.ajou.ac.kr/ns/board/notice.do |
| 자유전공학부 | https://pre.ajou.ac.kr/pre/community/notice.do |
| 전자공학과 | https://ece.ajou.ac.kr/ece/bachelor/notice.do |
| 정치외교학과 | https://pol.ajou.ac.kr/pol/activity/notice.do |
| 지능형반도체공학과 | https://aisemi.ajou.ac.kr/aisemi/board/notice.do |
| 첨단ICT융합대학 | https://it.ajou.ac.kr/it/board/notice.do |
| 첨단바이오융합대학 | https://ibio.ajou.ac.kr/ibio/data/data01.do |
| 첨단신소재공학과 | https://mse.ajou.ac.kr/mse/board/notice.do |
| 프런티어과학학부 | https://www.ajou.ac.kr/frontiers/board/notice.do |
| 행정학과 | https://pba.ajou.ac.kr/pba/activity/notice.do |
| 화학공학과 | https://che.ajou.ac.kr/che/board/notice.do |
| 화학과 | https://chem.ajou.ac.kr/chem/board/notice.do |
| 환경안전공학과 | https://env.ajou.ac.kr/env/board/notice.do |

## 수집 정보

각 공지 게시판에서 신규 공지를 감지하고, 공지 제목, 분류, 작성 부서, 본문, 이미지, 원문 URL 정보를 수집합니다.
