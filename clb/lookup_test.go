package clb

import (
	"fmt"
	"log"
	"testing"
)

var _ = fmt.Print // For debugging; delete when done.
var _ = log.Print // For debugging; delete when done.

func ExampleLookup() {
	server := DNSServer{Address: "8.8.8.8", Port: "53"}

	address, err := LookupAddress(server, "_xmpp-server._tcp.google.com")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", address)
	// Output: {Address:74.125.142.put:  Port:5269}
}

func Test_LookupGoogleXmppService_ReturnsAddress(t *testing.T) {
	// given
	port := "5269"
	srvName := "_xmpp-server._tcp.google.com"
	server := DNSServer{Address: "8.8.8.8", Port: "53"}

	// when
	address, err := LookupAddress(server, srvName)

	// then
	if err != nil {
		t.Error(err)
	}
	if port != address.Port {
		t.Errorf("port '%s' not expected", address.Port)
	}
	fmt.Printf("%+v", address)
}
