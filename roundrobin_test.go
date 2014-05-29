package main

import (
	"fmt"
	"github.com/benschw/consul-clb-go/roundrobinclb"
	"log"
	"testing"
)

var _ = fmt.Print // For debugging; delete when done.
var _ = log.Print // For debugging; delete when done.

func ExampleRoundRobinLookup() {
	srvName := "foo.service.fliglio.com"
	c := roundrobinclb.NewRoundRobinClb("127.0.0.1", "53")
	address, err := c.GetAddress(srvName)
	if err != nil {
		fmt.Print(err)
	}

	if address.Port == 8001 {
		fmt.Printf("%s", address)
	} else {
		address2, err := c.GetAddress(srvName)
		if err != nil {
			fmt.Print(err)
		}
		fmt.Printf("%s", address2)
	}
	// Output: 0.1.2.3:8001

}

//strconv.FormatInt(int64(srv.Port), 10)
func TestRoundRobinLookup(t *testing.T) {
	// given
	srvName := "foo.service.fliglio.com"
	c := roundrobinclb.NewRoundRobinClb("127.0.0.1", "53")

	// when
	address, err := c.GetAddress(srvName)
	// address2, err := c.GetAddress(srvName)

	// then
	if err != nil {
		t.Error(err)
	}

	if address.Port == 8001 && address.Address == "0.1.2.3" {
		return
	} else if address.Port == 8002 && address.Address == "4.5.6.7" {
		return
	} else {
		t.Errorf("port '%d' not expected with address: '%s'", address.Port, address.Address)
	}

}
