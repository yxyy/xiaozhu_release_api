FROM golang:1.23.3 AS builder

#设置代理
RUN go env -w GO111MODULE=on && go env -w  GOPROXY=https://goproxy.cn,direct

# 设置工作目录
WORKDIR /app/xiaozhu-backend

# 将当前项目目录的所有文件拷贝到容器的 WORKDIR 目录中
COPY . .

#下载依赖包
RUN  go mod tidy

# 静态编译 Go 项目
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o cmd/xiaozhu-backend cmd/main.go

#CMD ["./cmd/xiaozhu-backend"]

#运行阶段
FROM alpine:latest

WORKDIR /app/xiaozhu-backend/cmd

## 将编译后的二进制文件从构建阶段复制到运行阶段
COPY  --from=builder  /app/xiaozhu-backend/ /app/xiaozhu-backend/

# 运行
CMD ["./xiaozhu-backend"]