# Dockerfile for Post Service
FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build -o post-service ./cmd/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/post-service .
COPY .env .
EXPOSE 50052
CMD ["./post-service"]
