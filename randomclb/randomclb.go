package randomclb

import (
	"github.com/benschw/consul-clb-go/clb"
	"math/rand"
)

func NewRandomClb(address string, port string) clb.LoadBalancer {
	cfg := clb.DNSServerConfig{Address: address, Port: port}
	lb := RandomClb{serverConfig: cfg}
	return lb
}

type RandomClb struct {
	serverConfig clb.DNSServerConfig
}

func (lb RandomClb) GetAddress(name string) (clb.Address, error) {
	add := clb.Address{}

	srvs, err := clb.LookupSRV(lb.serverConfig, name)
	if err != nil {
		return add, err
	}
	//	log.Printf("%+v", srvs)

	srv := srvs[rand.Intn(len(srvs))]

	ip, err := clb.LookupA(lb.serverConfig, srv.Target)
	if err != nil {
		return add, err
	}

	return clb.Address{Address: ip, Port: srv.Port}, nil
}
