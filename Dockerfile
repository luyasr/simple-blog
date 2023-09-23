FROM golang:latest as builder
ENV GOPROXY https://goproxy.cn,direct
WORKDIR /app
COPY . /app/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o simple-blog .
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/simple-blog /app/simple-blog
RUN apk add --no-cache tzdata && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone && \
    chmod +x /app/simple-blog
EXPOSE 8080
CMD ["./simple-blog"]