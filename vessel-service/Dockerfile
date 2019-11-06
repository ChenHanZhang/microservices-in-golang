FROM golang:1.13 as builder

WORKDIR /go/src/github.com/ChenHanZhang/microservices-in-golang/vessel-service

COPY . .

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io

RUN go get
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/ChenHanZhang/microservices-in-golang/vessel-service/vessel-service .

CMD ["./vessel-service"]