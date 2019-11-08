FROM golang:1.13 as builder

# 将工作目录设置为当前微服务在gopath中的完整路径
WORKDIR /go/src/github.com/ChenHanZhang/microservices-in-golang/user-cli

# 将代码复制到工作目录中
COPY . .

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .


FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/ChenHanZhang/microservices-in-golang/user-cli/user-cli .

CMD ["./user-cli"]