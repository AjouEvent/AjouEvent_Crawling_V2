# syntax=docker/dockerfile:1.7

FROM golang:1.27-alpine AS dev
WORKDIR /workspace
RUN go install github.com/air-verse/air@latest
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod go mod download
COPY . .
CMD ["air"]
EXPOSE 1323

# 빌드 머신 아키텍처에서 교차 컴파일한다 (arm64 이미지도 에뮬레이션 없이 빌드)
FROM --platform=$BUILDPLATFORM golang:1.27-alpine AS build
ARG TARGETOS
ARG TARGETARCH
WORKDIR /workspace
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod go mod download
COPY . .
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=${TARGETOS:-linux} GOARCH=${TARGETARCH:-amd64} \
    go build -trimpath -ldflags="-s -w" -o crawling .

FROM alpine:3.23 AS prod
WORKDIR /workspace
RUN apk add --no-cache tzdata ca-certificates \
 && cp /usr/share/zoneinfo/Asia/Seoul /etc/localtime \
 && echo "Asia/Seoul" > /etc/timezone \
 && addgroup -S -g 10001 app && adduser -S -u 10001 -G app app
ENV TZ=Asia/Seoul
COPY --from=build /workspace/crawling .
COPY config config
USER app
CMD ["./crawling"]
