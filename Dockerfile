FROM golang:1.19-alpine as builder

RUN apk --no-cache add git

COPY . /root/build

WORKDIR /root/build

RUN GOPRIVATE=github.com/houko/wechatgpt GOPROXY=https://goproxy.cn,direct go mod download && go build -o server main.go

FROM alpine:latest as prod

RUN apk --no-cache add ca-certificates

WORKDIR /root/

RUN mkdir /root/data

COPY --from=0 /root/build/server .

CMD ["./server"]