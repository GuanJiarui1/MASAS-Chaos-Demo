# 第一阶段：构建
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o chaos-demo-app .

# 第二阶段：运行
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/chaos-demo-app .
EXPOSE 8080
# 关键修正：确保此处文件名与COPY指令来源的文件名一致
CMD ["./chaos-demo-app"]