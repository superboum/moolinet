FROM golang:alpine

RUN apk update && \
    apk add git && \
    adduser -S moolinet

COPY ./moolinet-write /usr/bin

WORKDIR /home/moolinet
USER moolinet
