FROM golang:1.24.2

WORKDIR /src/app

COPY . .

RUN go build -o main ./cmd/main.go

CMD ["./main"]