# syntax=docker/dockerfile:1.2

# 针对 x86_64 架构使用官方的 Golang 镜像作为基础镜像
FROM golang:1.20.8-alpine AS x86_64_builder

WORKDIR /app

COPY . .

RUN go build -o xapi-gateway main.go

# 使用 Alpine Linux 作为最终的基础镜像，针对 x86_64 架构
FROM alpine:latest AS x86_64_final

# 安装 GLIBC 和其他运行时库
RUN apk --no-cache add ca-certificates libc6-compat

WORKDIR /app

COPY --from=x86_64_builder /app/xapi-gateway .

COPY ./conf /app/conf

EXPOSE 8080

CMD ["./xapi-gateway"]

# 针对 ARM64 架构使用官方的 Golang 镜像作为基础镜像
FROM golang:1.20.8-alpine AS arm64v8_builder

WORKDIR /app

COPY . .

RUN go build -o xapi-gateway main.go

# 使用 Alpine Linux 作为最终的基础镜像，针对 ARM64 架构
FROM alpine:latest AS arm64v8_final

RUN apk --no-cache add ca-certificates libc6-compat

WORKDIR /app

COPY --from=arm64v8_builder /app/xapi-gateway .

COPY ./conf /app/conf

EXPOSE 8080

CMD ["./xapi-gateway"]