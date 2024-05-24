FROM golang:1.21 as builder

# 设置工作目录
WORKDIR /app

# 将go.mod和go.sum文件复制到工作目录，并下载依赖
COPY go.mod .
COPY go.sum .
RUN go mod download

# 将代码复制到工作目录
COPY . .

# 编译程序
RUN go build -o main .

# 使用轻量级的镜像
FROM alpine:latest

# 设置工作目录
WORKDIR /app

# 复制编译好的可执行文件
COPY --from=builder /app/main .

# 暴露监听端口
EXPOSE 8888

# 运行程序
CMD ["./main"]