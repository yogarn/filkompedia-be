FROM golang:1.23.6-alpine

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o main cmd/app/main.go

CMD ["./main"]
