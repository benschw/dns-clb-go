
all: build

deps:
	go get

build:
	go build -o demo

test:
	go test ./...

