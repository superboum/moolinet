#!/bin/bash
BASE=`dirname $0`
go build -o $BASE/moolinet-write $BASE/../tools/moolinet-write/main.go
docker build -t superboum/moolinet-golang $BASE -f $BASE/golang.dockerfile
