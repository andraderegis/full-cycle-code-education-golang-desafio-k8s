FROM golang:1.14-alpine as builder
RUN apk add --no-cache upx
WORKDIR /go/
COPY . ./src/desafio-k8s/go-k8s
RUN GOOS=linux go build -ldflags "-s -w" -tags netgo -a ./src/desafio-k8s/go-k8s/main.go
RUN upx --brute main

FROM busybox:latest
WORKDIR /go-k8s/
COPY --from=builder /go/main .
EXPOSE 8000
CMD ["./main"]