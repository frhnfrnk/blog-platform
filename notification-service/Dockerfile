# Dockerfile for Notification Service
FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build -o notification-service ./cmd/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/notification-service .
COPY .env .
EXPOSE 50054
CMD ["./notification-service"]
