FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build -o user-service ./cmd/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/user-service .
COPY .env .
EXPOSE 50051
CMD ["./user-service"]
