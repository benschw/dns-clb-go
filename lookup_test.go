package main

import (
	"fmt"
	"github.com/benschw/consul-clb-go/randomclb"
	"log"
	"testing"
)

var _ = fmt.Print // For debugging; delete when done.
var _ = log.Print // For debugging; delete when done.

func ExampleLookup() {
	c := randomclb.NewRandomClb("8.8.8.8", "53")
	address, err := c.GetAddress("_xmpp-server._tcp.google.com")
	//{Address:173.194.73.125:  Port:5269}

	if err != nil {
		panic(err)
	}

	fmt.Print(address.Port)
	// Output: 5269

}

//strconv.FormatInt(int64(srv.Port), 10)
func Test_LookupGoogleXmppService_ReturnsAddress(t *testing.T) {
	// given
	port := uint16(5269)
	srvName := "_xmpp-server._tcp.google.com"
	c := randomclb.NewRandomClb("8.8.8.8", "53")

	// when
	address, err := c.GetAddress(srvName)

	// then
	if err != nil {
		t.Error(err)
	}
	if port != address.Port {
		t.Errorf("port '%s' not expected", address.Port)
	}
}
