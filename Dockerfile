FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

RUN go install github.com/air-verse/air@latest

COPY . .

ENV PATH="/usr/local/go/bin:${PATH}"
ENTRYPOINT [ "air" ]
