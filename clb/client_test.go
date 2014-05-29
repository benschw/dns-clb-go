package clb

import (
	"fmt"
	"github.com/benschw/consul-clb-go/roundrobinclb"
	"log"
	"testing"
)

var _ = fmt.Print // For debugging; delete when done.
var _ = log.Print // For debugging; delete when done.

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
