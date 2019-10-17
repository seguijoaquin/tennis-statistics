FROM golang:alpine

WORKDIR /go/src/github.com/seguijoaquin/tennis-statistics/src/app

COPY ./app /go/src/github.com/seguijoaquin/tennis-statistics/src/app

RUN apk add git

RUN go get github.com/streadway/amqp

RUN go build

# CMD ./app