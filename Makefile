
all: build

deps:
	go get github.com/miekg/dns
	go get launchpad.net/goyaml

build:
	go build -o demo

test:
	go test -i github.com/benschw/consul-clb-go
	go test -v github.com/benschw/consul-clb-go

