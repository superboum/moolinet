# Source
# https://ariejan.net/2015/10/03/a-makefile-for-golang-cli-tools/

SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -regex '.*\.go\|.*\.html\|.*\.css\|.*\.js\|.*\.json')

BINARY=./release/moolinet-all

VERSION=v0.3
LDFLAGS=-ldflags "-X main.Version=${VERSION}"

docker=1.12
DOCKER_1.13=master
DOCKER_1.12=667315576fac663bd80bbada4364413692e57ac6

.DEFAULT_GOAL: $(BINARY)
$(BINARY): $(SOURCES)
	mkdir -p release/
	cp -r moolinet.json challenges/ static/ release/
	go build ${LDFLAGS} -o ${BINARY} moolinet-all/main.go

prepare:
	go get -d -v ./...
	cd ${GOPATH}/src/github.com/docker/docker && \
		echo "-> RESET" && git clean -fdx "" && \
		echo "-> MASTER" && git checkout master && \
		echo "-> PULL" && git pull && \
		echo "-> CHECKOUT" && git checkout ${DOCKER_${docker}}
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install

.PHONY: install
install: generate
	go install ${LDFLAGS} ./...

.PHONY: clean
clean:
	if [ -d release/ ] ; then rm -r release/ ; fi

.PHONY: test
test: generate
	go test ./...

.PHONY: lint
lint: install
	gometalinter -j 1 -t --deadline 100s \
		--exclude "Errors unhandled." \
		--exclude "moo.go" \
		--disable gotype --disable interfacer ./...

.PHONY: generate
generate:
	go generate ./...
