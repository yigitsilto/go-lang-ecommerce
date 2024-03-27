FROM golang:1.18.4-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o main

EXPOSE 3001

CMD ["./main"]
