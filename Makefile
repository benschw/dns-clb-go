
all: build

deps:
	go get github.com/miekg/dns

build:
	go build -o demo

