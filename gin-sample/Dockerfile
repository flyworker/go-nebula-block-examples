FROM golang:1.18-bullseye

ENV GO111MODULE=on

WORKDIR go-swan-examples
COPY . .

RUN go build -o main main.go
EXPOSE 8010
CMD ["./main"]

