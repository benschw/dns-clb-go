
all: build

deps:
	go get github.com/miekg/dns

build:
	go build -o demo

test:
	go test -i github.com/benschw/consul-clb-go/clb
	go test -v github.com/benschw/consul-clb-go/clb

