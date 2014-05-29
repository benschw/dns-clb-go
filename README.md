[![Build Status](https://drone.io/github.com/benschw/consul-clb-go/status.png)](https://drone.io/github.com/benschw/consul-clb-go/latest)

[![GoDoc](http://godoc.org/github.com/benschw/consul-clb-go?status.png)](http://godoc.org/github.com/benschw/consul-clb-go)

# Consul Client Load Balancer for Go

randomly selects a `SRV` record answer, then resolves its `A` record to an ip, and returns an `Address` structure:

	type Address struct {
		Address string
		Port    uint16
	}


example:
	

	srvName := "my-svc.service.consul"
	c := NewClb("127.0.0.1", "8600", RoundRobin)
	address, err := c.GetAddress(srvName)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", address.String())
	// Output: 0.1.2.3:8001

