FROM golang:alpine

RUN apk update && \
    apk add git && \
    adduser -S moolinet

WORKDIR /home/moolinet
USER moolinet
