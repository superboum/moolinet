moolinet
========

![Travis Moolinet](https://api.travis-ci.org/superboum/moolinet.svg?branch=master)

## Requirements

 * Go 1.7
 * Docker 1.12.x (API 24) up and running on your system


## Installation

```
go get -d github.com/superboum/moolinet/...
go install github.com/superboum/moolinet/...
git --git-dir ../../docker/docker/.git checkout 667315576fac663bd80bbada4364413692e57ac6
go test -v ./... # (as root)
```

*I'm considering using golang vendors to fix problems linked to the docker API change*

## Images

```
sudo docker build -t superboum/moolinet-golang:v1 ./images -f ./images/golang.dockerfile
```

## Run the tests

Tests must be run as root because we are making some calls to Docker. You'll probably have to define your `GOPATH`. An example:

```
sudo su -c "GOPATH=/your/go/path go test github.com/superboum/moolinet/..."
```
