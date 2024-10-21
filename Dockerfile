FROM golang:1.23.2

ENV TODO_PORT=7540
ENV TODO_PASSWORD=1234
ENV TODO_DBFILE=scheduler.db
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

FROM ubuntu:latest

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /todoserver

EXPOSE ${TODO_PORT}
CMD ["./todoapp"]