FROM golang:1.22

RUN mkdir /go/src/myapp
WORKDIR /go/src/myapp

RUN go install github.com/cosmtrek/air@v1.51.0
