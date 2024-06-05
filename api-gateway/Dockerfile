FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o /api-gateway ./cmd/main.go

EXPOSE 8080

CMD ["/api-gateway"]
