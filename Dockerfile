FROM golang:1.22.8  as builder
ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct
RUN cd / && mkdir app
WORKDIR /app
COPY ./go.mod .
COPY ./go.sum .
RUN go mod download
COPY . /app
RUN  go build -ldflags '-s -w' -o ./build/mm-market-server main.go

# =====
FROM alpine:3.19.1

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/cadok-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

WORKDIR /app

COPY --from=builder /app/build/mm-market-server /app/mm-market-server
VOLUME /app
CMD ["/app/mm-market-server","--config","./config.json"]