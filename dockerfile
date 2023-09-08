# 使用官方的 Golang 镜像作为基础镜像
FROM golang:1.20 AS builder

WORKDIR /app

COPY . .

RUN go build -o xapi-gateway main.go

# 使用 Alpine Linux 作为最终的基础镜像
FROM alpine:latest

# 安装 GLIBC 和其他运行时库
RUN apk --no-cache add ca-certificates libc6-compat

WORKDIR /app

COPY --from=builder /app/xapi-gateway .

COPY ./conf /app/conf

EXPOSE 8080

CMD ["./xapi-gateway"]