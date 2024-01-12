FROM golang:1.20-alpine3.16

LABEL maintainer="raudra.gmail@nitt.com"

RUN apk add build-base bash git


WORKDIR /rule-manager

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . /rule-manager

RUN go mod tidy

RUN go build -o /usr/local/bin/rule-manager

CMD rule-manager


