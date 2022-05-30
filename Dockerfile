FROM golang:1.14.9-alpine AS builder
RUN go build -o app -v .

CMD ["./app"]