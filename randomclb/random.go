package randomclb

import (
	"fmt"
	"github.com/benschw/consul-clb-go/clb"
	"math/rand"
)

func NewRandomClb(address string, port string) clb.LoadBalancer {
	lb := RandomClb{serverStr: fmt.Sprintf("%s:%s", address, port)}
	return lb
}

type RandomClb struct {
	serverStr string
}

func (lb RandomClb) GetAddress(name string) (clb.Address, error) {
	add := clb.Address{}

	srvs, err := clb.LookupSRV(lb.serverStr, name)
	if err != nil {
		return add, err
	}
	//	log.Printf("%+v", srvs)

	srv := srvs[rand.Intn(len(srvs))]

	ip, err := clb.LookupA(lb.serverStr, srv.Target)
	if err != nil {
		return add, err
	}

	return clb.Address{Address: ip, Port: srv.Port}, nil
}
