package clb

import (
	"fmt"
	"log"
	"testing"
)

var _ = fmt.Print // For debugging; delete when done.
var _ = log.Print // For debugging; delete when done.

// func Test_LookupConsul_ReturnsConsulAddress(t *testing.T) {
// 	// given
// 	port := "8300"

// 	// when
// 	srvName := "consul.service.consul"
// 	address, err := LookupAddress(srvName)

// 	// then
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	if port != address.Port {
// 		t.Errorf("port %s not expected", address.Port)
// 	}
// 	fmt.Printf("%+v", address)
// }

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
