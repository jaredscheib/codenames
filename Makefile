VERSION=0.0.0
COMMIT=$(shell git rev-parse --short=8 HEAD)
DATE=$(shell date +%Y%m%d-%H:%M:%S)

SOURCES := $(shell find . -name '*.go' )

unexport LDFLAGS
LDFLAGS=-ldflags "-s -X main.version=${VERSION} -X main.commit=${COMMIT} -X main.date=${DATE}"
BINARY=codenames

.DEFAULT_GOAL := all

success:
	@echo "Codenames built successfully!"

all: build success

build: ${BINARY}

${BINARY}: $(SOURCES)
	go build -o ./bin/${BINARY} ${LDFLAGS} ./cmd/*.go