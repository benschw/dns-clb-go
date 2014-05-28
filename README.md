<!-- [![Build Status](https://drone.io/github.com/BancVue/a10mgt/status.png)](https://drone.io/github.com/BancVue/a10mgt/latest) -->

[![Build Status](https://travis-ci.org/benschw/consul-clb-go.png?branch=master)](https://travis-ci.org/benschw/consul-clb-go)

[![GoDoc](http://godoc.org/github.com/benschw/consul-clb-go?status.png)](http://godoc.org/github.com/benschw/consul-clb-go)

# Consul Client Load Balancer for Go

randomly selects a `SRV` record answer, then resolves its `A` record to an ip, and returns an `Address` structure:

	type Address struct {
		Address string
		Port string
	}


example:
	
	svcName := "my-svc"

	srvRecord := svcName + ".service.consul"
	address, err := clb.LookupAddress(srvRecord)
	if err != nil {
		panic(err)
	}

	fmt.Print(address.Address + ":" + address.Port)

