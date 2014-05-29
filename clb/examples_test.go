package clb

import (
	"fmt"
)

func ExampleRoundRobinLookup() {
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
