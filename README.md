[![Build Status](https://drone.io/github.com/benschw/consul-clb-go/status.png)](https://drone.io/github.com/benschw/consul-clb-go/latest)
[![GoDoc](http://godoc.org/github.com/benschw/consul-clb-go?status.png)](http://godoc.org/github.com/benschw/consul-clb-go)

# Consul Client Load Balancer for Go

Selects a `SRV` record answer according to specified load balancer algorithm, then resolves its `A` record to an ip, and returns an `Address` structure:

	type Address struct {
		Address string
		Port    uint16
	}


## Example:
	

	srvName := "my-svc.service.consul"
	c := clb.NewClb("127.0.0.1", "8600", clb.RoundRobin)
	address, err := c.GetAddress(srvName)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", address.String())
	// Output: 0.1.2.3:8001

## Development
tests are run against some fixture dns entries I set up on fliglio.com (`dig foo.service.fliglio.com SRV`).


- `make deps` install deps
- `make test` run all tests