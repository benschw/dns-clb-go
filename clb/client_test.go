package clb

import (
	"fmt"
	"github.com/benschw/consul-clb-go/roundrobinclb"
	"log"
	"testing"
)

var _ = fmt.Print // For debugging; delete when done.
var _ = log.Print // For debugging; delete when done.

func ExampleRoundRobin() {
	srvName := "foo.service.fliglio.com"
	c := NewRoundRobinClb("8.8.8.8", "53")
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

func doStuff(c LoadBalancer) error {
	srvName := "foo.service.fliglio.com"
	_, err := c.GetAddress(srvName)
	return err
}

// this is a rediculous test, but i got confused by the interface
func TestLoadBalancerInterface(t *testing.T) {
	// given
	c := roundrobinclb.NewRoundRobinClb("8.8.8.8", "53")

	// when
	err := doStuff(c)

	// then
	if err != nil {
		t.Error(err)
	}
}

func TestRoundRobinFacade(t *testing.T) {
	//given
	c := NewClb("8.8.8.8", "53", RoundRobin)

	// when
	err := doStuff(c)

	// then
	if err != nil {
		t.Error(err)
	}
}
func TestRandomFacade(t *testing.T) {
	//given
	c := NewClb("8.8.8.8", "53", Random)

	// when
	err := doStuff(c)

	// then
	if err != nil {
		t.Error(err)
	}
}
