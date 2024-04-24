FROM golang:1.18.4-alpine

WORKDIR /app

ARG PORT
ENV PORT=${WORKER_PORT}

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o main

EXPOSE $PORT

CMD ["./main"]
