package Dockerfile

FROM golang:1.22

RUN mkdir /cmd
WORKDIR /cmd

COPY ./* ./

RUN go build -o app

ENTRYPOINT ["/cmd/app"]