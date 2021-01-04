FROM golang:1.14-alpine as builder
RUN apk add --no-cache upx
WORKDIR /go/src/go-k8s
COPY . .
RUN GOOS=linux go build -ldflags "-s -w" -tags netgo -a main.go
RUN upx --brute main

FROM busybox:latest
WORKDIR /go-k8s
COPY --from=builder /go/src/go-k8s/main .
EXPOSE 8000
ENTRYPOINT ["./main"]