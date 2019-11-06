FROM golang:1.13 as builder

WORKDIR /go/src/github.com/ChenHanZhang/microservices-in-golang/consignment-cli

COPY . .

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn

RUN go get -u
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .


FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY consignment.json /app/consignment.json
COPY --from=builder /go/src/github.com/ChenHanZhang/microservices-in-golang/consignment-cli/consignment-cli .

CMD ["./consignment-cli"]