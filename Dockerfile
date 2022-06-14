# syntax=docker/dockerfile:1
FROM golang:1.18
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /docker-go-cards

EXPOSE 8080

CMD [ "/docker-go-cards" ]