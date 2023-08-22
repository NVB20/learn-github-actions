FROM golang:1.20.7-alpine3.18

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY  ./go-app/go_log_test.go .
COPY ./go-app/go_log.go .

RUN go build go_log.go

CMD ["./go_log"]