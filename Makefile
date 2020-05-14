.PHONY : build run fresh test clean pack-releases

BIN := kubekutr

HASH := $(shell git rev-parse --short HEAD)
COMMIT_DATE := $(shell git show -s --format=%ci ${HASH})
BUILD_DATE := $(shell date '+%Y-%m-%d %H:%M:%S')
VERSION := ${HASH} (${COMMIT_DATE})

STATIC := ./templates:/templates

deps:
	go get -u github.com/knadh/stuffbin/...

build:
	go build -o ${BIN} -ldflags="-X 'main.buildVersion=${VERSION}' -X 'main.buildDate=${BUILD_DATE}'"
	stuffbin -a stuff -in ${BIN} -out ${BIN} ${STATIC}

test-local: fresh
	./${BIN} -c config.yml s

validate-local: fresh
	./${BIN} -c config.yml s | kubeval --strict

run:
	./${BIN}

fresh: clean build

test:
	go test

clean:
	go clean
	- rm -f ${BIN}

# pack-releases runns stuffbin packing on a given list of
# binaries. This is used with goreleaser for packing
# release builds for cross-build targets.
pack-releases:
	$(foreach var,$(RELEASE_BUILDS),stuffbin -a stuff -in ${var} -out ${var} ${STATIC} $(var);)
