SOURCES := $(shell find . -name '*.go' )

BINARY=codenames

.DEFAULT_GOAL := all

all: build

build: ${BINARY}

${BINARY}: $(SOURCES)
	go build -o ./bin/${BINARY} ./main.go