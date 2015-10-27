[![Build Status](https://drone.io/github.com/benschw/dns-clb-go/status.png)](https://drone.io/github.com/benschw/dns-clb-go/latest)
[![GoDoc](http://godoc.org/github.com/benschw/dns-clb-go?status.png)](http://godoc.org/github.com/benschw/dns-clb-go)


_Deprecated in favor of [srv-lb](https://github.com/benschw/srv-lb), a rewrite with a cleaner interface._


# DNS Client Load Balancer for Go

Selects a `SRV` record answer according to specified load balancer algorithm, then resolves its `A` record to an ip, and returns an `Address` structure:

	type Address struct {
		Address string
		Port    uint16
	}


## Example:
	
	// uses dns server configured in /etc/resolv.conf
	srvName := "my-svc.service.consul"
	c := clb.New() 
	address, err := c.GetAddress(srvName)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", address.String())
	// Output: 0.1.2.3:8001

### or configure explicitely

	srvName := "my-svc.service.consul"
	c := clb.NewClb("127.0.0.1", "8600", clb.RoundRobin)
	address, err := c.GetAddress(srvName)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", address.String())
	// Output: 0.1.2.3:8001

### or use an `AddressProvider`
	
	// uses dns server configured in /etc/resolv.conf
	ap := NewAddressProvider("my-svc.service.consul")
	address, err := ap.GetAddress()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", address.String())
	// Output: 0.1.2.3:8001


## Development
tests are run against some fixture dns entries I set up on fligl.io (`dig foo.service.fligl.io SRV`).


- `make deps` install deps
- `make test` run all tests

## Notes for Consul / Confd Cluster Demo
This is a fork from the original project https://github.com/benschw/consul-clb-go which contains the `demo` service used in my blog post outlining how to use Consul for service discovery and configuration management while using Confd and DNS to keep your applications decoupled from the specifics of Consul.

- download `demo` service here: https://github.com/benschw/consul-clb-go/releases/tag/v0.1.0
- blog post outlining the demo: http://txt.fliglio.com/2014/05/encapsulated-services-with-consul-and-confd/


