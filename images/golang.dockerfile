FROM golang:alpine

RUN apk update && \
    apk add git openssl && \
    adduser -S moolinet

RUN go get -u github.com/alecthomas/gometalinter && \
    gometalinter --install && \
    chmod -R 777 /go

COPY ./moolinet-write /usr/bin

WORKDIR /home/moolinet
USER moolinet
