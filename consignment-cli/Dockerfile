FROM golang:1.13 as builder

WORKDIR /go/src/github.com/ChenHanZhang/microservices-in-golang/consignment-cli

COPY . .

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY consignment.json /app/consignment.json
COPY --from=builder /go/src/github.com/ChenHanZhang/microservices-in-golang/consignment-cli/consignment-cli .

ENTRYPOINT ["./consignment-cli", "consignment.json", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7ImlkIjoiZWU0ZTVhNWItNzk2ZS00NGYxLWE3ZTYtNTdkYzU4M2ZmZGY2IiwibmFtZSI6InRmaG5pYSIsImNvbXBhbnkiOiJCVVBUIiwiZW1haWwiOiJ0ZmhuaWFAMTYzLmNvbSIsInBhc3N3b3JkIjoiJDJhJDEwJGFLbnJ5SjJUWFdLQ1BKVWEvLnRhVy5sajVaM1p3akl3ekxVc3AyYW9FeGM4OTBpR3hzZHR1In0sImV4cCI6MTU3MzQ2MzQ5OSwiaXNzIjoiZ28ubWljcm8uc3J2LnVzZXIifQ.VJR042fZjfadbcF2OlIBtAipYNQuIlfLEOJtW8mbOG8"]
CMD ["./consignment-cli"]