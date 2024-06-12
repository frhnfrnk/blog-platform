# Dockerfile for API Gateway
FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build -o api-gateway ./cmd/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/api-gateway .
COPY .env .
EXPOSE 8080
CMD ["./api-gateway"]
