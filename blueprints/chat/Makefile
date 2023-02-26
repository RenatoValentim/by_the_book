.PHONY: all deps run test

all: deps run

deps:
	@go mod tidy

run:
	@go run *.go

test:
	@go test -v -cover -failfast ./...
