
all: build

deps:
	go get github.com/miekg/dns
	go get launchpad.net/goyaml

build:
	go build -o demo

test:
	go test ./...

