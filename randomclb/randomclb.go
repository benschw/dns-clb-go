package randomclb

import (
	"github.com/benschw/consul-clb-go/clb"
	"math/rand"
)

func NewRandomClb(address string, port string) clb.LoadBalancer {
	c := clb.NewDNSClient(address, port)
	lb := RandomClb{client: c}
	return lb
}

type RandomClb struct {
	client *clb.DNSClient
}

func (lb RandomClb) GetAddress(name string) (clb.Address, error) {
	add := clb.Address{}

	srvs, err := lb.client.LookupSRV(name)
	if err != nil {
		return add, err
	}
	//	log.Printf("%+v", srvs)

	srv := srvs[rand.Intn(len(srvs))]

	aAdd, err := lb.client.LookupA(srv.Address)
	if err != nil {
		return add, err
	}

	return clb.Address{Address: aAdd.Address, Port: srv.Port}, nil
}
