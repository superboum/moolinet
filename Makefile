# Source
# https://ariejan.net/2015/10/03/a-makefile-for-golang-cli-tools/

# MOOLINET-ALL
SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -regex '.*\.go\|.*\.html\|.*\.css\|.*\.js\|.*\.json')
BINARY=./release/moolinet-all

# MOOLINET-FUZZ
MOOLINET_FUZZ=./release/tools/moolinet-fuzz
FUZZ_SOURCEDIR=./tools/moolinet-fuzz
FUZZ_SOURCES := $(shell find $(FUZZ_SOURCEDIR) -regex '.*\.go\|.*\.y')

# MOOLINET-WRITE
MOOLINET_WRITE=./release/tools/moolinet-write
WRITE_SOURCEDIR=./tools/moolinet-write
WRITE_SOURCES := $(shell find $(WRITE_SOURCEDIR) -regex '.*\.go')

VERSION=v0.3
LDFLAGS=-ldflags "-X main.Version=${VERSION}"

docker=master
DOCKER_master=master
DOCKER_1.12=667315576fac663bd80bbada4364413692e57ac6

.DEFAULT_GOAL: release

release: $(BINARY) $(MOOLINET_FUZZ) $(MOOLINET_WRITE)

$(BINARY): $(SOURCES)
	mkdir -p release/
	cp -r moolinet.json challenges/ static/ release/
	go build ${LDFLAGS} -o ${BINARY} moolinet-all/main.go

$(MOOLINET_FUZZ): $(FUZZ_SOURCES) generate
	mkdir -p release/tools
	go build ${LDFLAGS} -o ${MOOLINET_FUZZ} ./tools/moolinet-fuzz/main.go

$(MOOLINET_WRITE): $(WRITE_SOURCES)
	mkdir -p release/tools
	go build ${LDFLAGS} -o ${MOOLINET_WRITE} ./tools/moolinet-write/main.go

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
		--exclude "tools/moolinet-fuzz/testdata" \
		--exclude "moo.go" \
		--disable gotype --disable interfacer ./...

.PHONY: generate
generate:
	go generate ./...
