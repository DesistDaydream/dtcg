FROM golang:1.18-alpine as builder
WORKDIR /root/prometheus-instrumenting

ENV CGO_ENABLED=0 \
    GO111MODULE=on \
    GOPROXY=https://goproxy.cn,https://goproxy.io,direct
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY ./ /root/prometheus-instrumenting
RUN go build -o jihuanshe-exporter ./cmd/jihuanshe_exporter/*.go

FROM alpine
# org.opencontainers.image.source 用于为 GitHub Package 提供标识符，以识别该镜像应该属于哪个仓库
LABEL org.opencontainers.image.source https://github.com/DesistDaydream/jihuanshe-exporter
WORKDIR /root/prometheus-instrumenting
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk update && \
    apk add --no-cache tzdata && \
    ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
ENV TZ=Asia/Shanghai
COPY --from=builder /root/prometheus-instrumenting/jihuanshe-exporter /usr/local/bin/jihuanshe-exporter
COPY ./cards /root/prometheus-instrumenting/cards
ENTRYPOINT  [ "/usr/local/bin/jihuanshe-exporter" ]