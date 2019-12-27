.PHONY : build run fresh test clean

BIN := kubekutr

HASH := $(shell git rev-parse --short HEAD)
COMMIT_DATE := $(shell git show -s --format=%ci ${HASH})
BUILD_DATE := $(shell date '+%Y-%m-%d %H:%M:%S')
VERSION := ${HASH} (${COMMIT_DATE})

STATIC := ./templates:/templates

build:
	go build -o ${BIN} -ldflags="-X 'main.buildVersion=${VERSION}' -X 'main.buildDate=${BUILD_DATE}'"
	stuffbin -a stuff -in ${BIN} -out ${BIN} ${STATIC}

run:
	./kubekutr

fresh: clean build run

test:
	go test

clean:
	go clean
	- rm -f ${BIN}
