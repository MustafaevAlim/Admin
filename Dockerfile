FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd/project/main.go

EXPOSE 8080

CMD ["/app/main"]