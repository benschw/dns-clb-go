package clb

import (
	"fmt"
	"github.com/benschw/dns-clb-go/dns"
	"log"
	"testing"
)

var _ = fmt.Print // For debugging; delete when done.
var _ = log.Print // For debugging; delete when done.

// Example with direct usage
func ExampleNewRoundRobinClb() {
	srvName := "foo.service.fliglio.com"
	lib := dns.NewLookupLib("8.8.8.8:53")
	c := NewRoundRobinClb(lib)
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

// Example with factory
func ExampleNewClb() {
	srvName := "foo.service.fliglio.com"
	c := NewClb("8.8.8.8", "53", RoundRobin)
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

func TestRoundRobinFacade(t *testing.T) {
	//given
	c := NewClb("8.8.8.8", "53", RoundRobin)

	// when
	srvName := "foo.service.fliglio.com"
	_, err := c.GetAddress(srvName)

	// then
	if err != nil {
		t.Error(err)
	}
}

func TestRandomFacade(t *testing.T) {
	//given
	c := NewClb("8.8.8.8", "53", Random)

	// when
	srvName := "foo.service.fliglio.com"
	_, err := c.GetAddress(srvName)

	// then
	if err != nil {
		t.Error(err)
	}
}

func TestTtlCacheFacade(t *testing.T) {
	//given
	c := NewTtlCacheClb("8.8.8.8", "53", Random, 5)

	// when
	srvName := "foo.service.fliglio.com"
	_, err := c.GetAddress(srvName)

	// then
	if err != nil {
		t.Error(err)
	}
}
