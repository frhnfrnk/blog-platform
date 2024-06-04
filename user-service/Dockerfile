FROM golang:1.19

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /main ./cmd/main.go

EXPOSE 50051

CMD ["/main"]
