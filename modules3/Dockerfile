FROM golang:1.17  as  builder

# docker中的工作目录
WORKDIR /build

ENV CGO_ENABLED=0
ENV GO111MODULE=off
ENV GOPROXY=https://goproxy.cn,direct
# 将当前目录同步到docker工作目录下，也可以只配置需要的目录和文件（配置目录、编译后的程序等）
COPY . ./

# 指定编译完成后的文件名，可以不设置使用默认的，最后一步要执行该文件名
RUN go build -o httpserver .
FROM scratch
COPY --from=builder /build/httpserver /
ENTRYPOINT ["/httpserver"]