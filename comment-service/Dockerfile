FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./
COPY cmd/ cmd/
COPY internal/ internal/

RUN go mod download

RUN go build -o comment-service ./cmd/main.go

EXPOSE 50053

CMD ["./comment-service"]
