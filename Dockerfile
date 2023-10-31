FROM golang:1.21

WORKDIR /app
COPY go.mod go.sum

RUN go install github.com/cosmtrek/air@latest

COPY . .

CMD "air"