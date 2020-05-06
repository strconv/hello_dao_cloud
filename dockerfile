# 基础镜像
FROM golang:latest as builder
# go mod
ENV GO111MODULE on
ENV CGO_ENABLED=0
ENV GOPROXY https://goproxy.io

# 工作目录
WORKDIR /app
COPY . .
# 这里将Golang依赖定义相关文件的copy放到最前面
COPY go.mod go.sum ./
RUN go mod download

# 编译 （注意 hello_dao_cloud 在 docker 容器里编译，并没有在宿主机现场编译）
RUN go build -o main .

FROM alpine:latest
WORKDIR /app/
# alpine镜像没有ca-certificates，需要进行安装
RUN apk update && apk --no-cache add ca-certificates
# 从builder stage的镜像里将二进制文件copy过来
COPY --from=builder /app/main .

# 声明运行时容器提供服务端口
EXPOSE 8082
CMD ["./main"]