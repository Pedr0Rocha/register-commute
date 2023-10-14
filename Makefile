all: test build

build:
	go build -o bin/register-commute cmd/main.go

test:
	go test -v ./...

clean:
	rm -f bin/register-commute

run: build
	./bin/register-commute

fmt:
	go fmt ./...

.PHONY: all build test clean run fmt
