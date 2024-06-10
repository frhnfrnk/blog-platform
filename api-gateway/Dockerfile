FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./
COPY cmd/ cmd/
COPY internal/ internal/

RUN go mod download

RUN go build -o api-gateway ./cmd/main.go

EXPOSE 8080

CMD ["./api-gateway"]
