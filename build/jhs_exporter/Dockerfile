FROM golang:1.20-alpine as builder
WORKDIR /dtcg

# 添加 cgo 环境
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk add build-base

ENV CGO_ENABLED=1 \
    GO111MODULE=on \
    GOPROXY=https://goproxy.cn,https://goproxy.io,direct
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY ./ /dtcg
RUN go build -o jhs-exporter ./cmd/jhs_exporter/*.go

FROM alpine
# org.opencontainers.image.source 用于为 GitHub Package 提供标识符，以识别该镜像应该属于哪个仓库
LABEL org.opencontainers.image.source https://github.com/DesistDaydream/dtcg
WORKDIR /dtcg
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk update && \
    apk add --no-cache tzdata && \
    ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
ENV TZ=Asia/Shanghai
COPY --from=builder /dtcg/jhs-exporter /dtcg/jhs-exporter
ENTRYPOINT  [ "/dtcg/jhs-exporter" ]