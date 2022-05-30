FROM golang:1.14.9-alpine AS builder
RUN mkdir /build
ADD go.mod main.go /build/
WORKDIR /build
RUN go build -o app -v .

FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/main /app/
WORKDIR /app
CMD ["./app"]