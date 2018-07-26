SOURCES := $(shell find . -name '*.go' )

BINARY=codenames

.DEFAULT_GOAL := all

success:
	@echo "Codenames built successfully!"

all: build success

build: ${BINARY}

${BINARY}: $(SOURCES)
	go build -o ./bin/${BINARY} ./main.go