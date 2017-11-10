moolinet
========

[![Build Status](http://ci.deuxfleurs.fr/job/moolinet/job/master/badge/icon)](http://ci.deuxfleurs.fr/job/moolinet/job/master/)

## Tested with

 * go1.9.2 linux/amd64
 * 17.05.0-ce (API 1.29)
 * Debian Testing

## Installation

```bash
# Download the project with its dependencies
go get -d github.com/superboum/moolinet/...

# Compile it
cd $GOPATH/src/github.com/superboum/moolinet
go install ./...

# Run it
moolinet-all -config moolinet.json

# Test it
go test -v github.com/superboum/moolinet/... # (you should be in the docker group or run this test as root)
```

You might need to checkout a specific tag/commit for the Docker Client, for example:

```
# Docker does not follow the go convention "don't break your API"
cd $GOPATH/src/github.com/docker/docker
git checkout 667315576fac663bd80bbada4364413692e57ac6
```

## Create a release

```bash
make release
```

You'll have everything you need to run moolinet, included its tool in the folder release

## Images

```
bash ./images/make.sh
```

