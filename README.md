moolinet
========

## Requirements

 * Go 1.7
 * Docker up and running on your system


## Installation

```
go get -d github.com/superboum/moolinet/...
go install github.com/superboum/moolinet/...
sudo ./moolinet-worker
```

## Images

```
sudo docker build -t superboum/moolinet-golang:v1 ./images -f ./images/golang.dockerfile
```

## Run the tests

Tests must be run as root because we are making some calls to Docker. You'll probably have to define your `GOPATH`. An example:

```
sudo su -c "GOPATH=/your/go/path go test github.com/superboum/moolinet/..."
```
