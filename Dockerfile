FROM golang:1.18.3-stretch AS builder

# 环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY="https://goproxy.cn,direct"

# 切换到工作目录
WORKDIR /build

# 将代码复制到容器中
COPY . .

# 下载依赖
RUN go mod download \
    && go get shershon1991/fund-analye-system/service/crawl/fund \
    && go get github.com/gin-gonic/gin/binding@v1.7.2

# 编译成二进制文件,二进制文件名：app
RUN go build -o app .


### --------- 二阶段，构建一个小镜像 ---------
FROM shershon/go-env:stretch-slim

# 项目目录
WORKDIR /www

# 从builder镜像中把二进制文件/build/app 拷贝到当前目录
COPY --from=builder /build/app /www
