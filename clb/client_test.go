package clb

import (
	"fmt"
	"log"
	"testing"

	"github.com/benschw/dns-clb-go/dns"
)

var _ = fmt.Print // For debugging; delete when done.
var _ = log.Print // For debugging; delete when done.

// Example with direct usage
func ExampleNewRoundRobinClb() {
	srvName := "foo.service.fligl.io"
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

// Example load balancer with default dns server
func ExampleNew() {
	c := New()
	address, err := c.GetAddress("foo.service.fligl.io")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", address.String())
	// Output: 0.1.2.3:8001
}

// Example address provider using defaults
func ExampleAddressProvider() {
	ap := NewAddressProvider("foo.service.fligl.io")
	address, err := ap.GetAddress()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", address.String())
	// Output: 0.1.2.3:8001
}

// Example with factory
func ExampleNewClb() {
	srvName := "foo.service.fligl.io"
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

func TestAddressProvider(t *testing.T) {
	// given
	c := NewClb("8.8.8.8", "53", RoundRobin)
	ap := &SRVAddressProvider{Lb: c, Address: "foo.service.fligl.io"}

	// when
	add, err := ap.GetAddress()

	// then
	if err != nil {
		t.Error(err)
	}

	if add.Port == 8001 && add.Address == "0.1.2.3" {
		return
	} else if add.Port == 8002 && add.Address == "4.5.6.7" {
		return
	} else {
		t.Errorf("address looks wrong: %+v", add)
	}

}

func TestRoundRobinFacade(t *testing.T) {
	//given
	c := NewClb("8.8.8.8", "53", RoundRobin)

	// when
	srvName := "foo.service.fligl.io"
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
	srvName := "foo.service.fligl.io"
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
	srvName := "foo.service.fligl.io"
	_, err := c.GetAddress(srvName)

	// then
	if err != nil {
		t.Error(err)
	}
}
