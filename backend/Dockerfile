# syntax=docker/dockerfile:1

FROM golang:1.17-alpine

RUN apk update && apk add --no-cache gcc && apk add --no-cache g++

WORKDIR /app

COPY . .

RUN go get -d -v ./...
RUN go install github.com/mattn/go-sqlite3 
RUN go install -v ./...

EXPOSE 8080

CMD [ "go", "run", "main.go" ]