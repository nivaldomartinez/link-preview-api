# syntax=docker/dockerfile:1

FROM golang:1.20.1-bullseye

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /m

EXPOSE 8080

CMD [ "/m" ]
