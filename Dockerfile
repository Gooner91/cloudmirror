FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

RUN go install github.com/air-verse/air@latest

COPY . .

ENTRYPOINT [ "air" ]
