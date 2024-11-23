FROM golang:1.23.2-alpine

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o main .

EXPOSE 8000

CMD ["./main"]
