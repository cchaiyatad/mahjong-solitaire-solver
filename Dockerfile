FROM golang:1.17 AS builder
WORKDIR /go/src/mss
COPY . .
RUN  go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app ./cmd/mss/main.go

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/mss/app ./bin/mss/
COPY --from=builder /go/src/mss/assets ./assets

WORKDIR /root/bin/mss
ENTRYPOINT ["./app"]  
