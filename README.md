# Consul CLB for Go


example:
	
	svcName := "my-svc"

	srvRecord := svcName + ".service.consul"
	address, err := clb.LookupAddress(srvRecord)
	if err != nil {
		panic(err)
	}

	fmt.Print(address.Address + ":" + address.Port)

