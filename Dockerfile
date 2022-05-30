FROM golang:1.14.9-alpine AS builder
RUN mkdir /build
ADD go.mod main.go /build/
WORKDIR /build
RUN go build -o app -v .

CMD ["./app"]