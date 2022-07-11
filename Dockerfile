FROM golang:1.18

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build cmd/main.go

CMD ["./main"]
