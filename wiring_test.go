package main

import (
	"fmt"
	"github.com/benschw/consul-clb-go/clb"
	"github.com/benschw/consul-clb-go/roundrobinclb"
	"log"
	"testing"
)

var _ = fmt.Print // For debugging; delete when done.
var _ = log.Print // For debugging; delete when done.

func doStuff(c clb.LoadBalancer) error {
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
	c := clb.NewClb("8.8.8.8", "53", clb.RoundRobin)

	// when
	err := doStuff(c)

	// then
	if err != nil {
		t.Error(err)
	}
}
func TestRandomFacade(t *testing.T) {
	//given
	c := clb.NewClb("8.8.8.8", "53", clb.Random)

	// when
	err := doStuff(c)

	// then
	if err != nil {
		t.Error(err)
	}
}
