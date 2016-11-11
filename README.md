moolinet
========

[![Build Status](http://ci.deuxfleurs.fr/job/moolinet/job/master/badge/icon)](http://ci.deuxfleurs.fr/job/moolinet/job/master/)

## Requirements

 * Go 1.7
 * Docker 1.12.x (API 24) up and running on your system


## Installation

```bash
# Download the project with its dependencies
go get -d github.com/superboum/moolinet/...

# Docker does not follow the go convention "don't break your API"
cd ../../docker/docker
git --git-dir ../../docker/docker/.git checkout 667315576fac663bd80bbada4364413692e57ac6
cd -

# Test it !
go test -v ./... # (you should be in the docker group or run this test as root)
```

## Images

```
sudo docker build -t superboum/moolinet-golang:v1 ./images -f ./images/golang.dockerfile
```

