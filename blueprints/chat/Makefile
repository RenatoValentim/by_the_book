.PHONY: all deps clean build chat
BINARY=./chat

all: deps run

deps:
	@go mod tidy

build:
	@go build

run: build
	@./chat

clean:
	@rm -rf ./chat
