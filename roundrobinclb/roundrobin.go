package roundrobinclb

import (
	"fmt"
	"github.com/benschw/consul-clb-go/clb"
)

func NewRoundRobinClb(address string, port string) clb.LoadBalancer {
	lb := RoundRobinClb{serverStr: fmt.Sprintf("%s:%s", address, port), i: 0}
	return lb
}

type RoundRobinClb struct {
	serverStr string
	i         int
}

func (lb RoundRobinClb) GetAddress(name string) (clb.Address, error) {
	add := clb.Address{}

	srvs, err := clb.LookupSRV(lb.serverStr, name)
	if err != nil {
		return add, err
	}
	//	log.Printf("%+v", srvs)
	if len(srvs)-1 > lb.i {
		lb.i = 0
	}
	srv := srvs[lb.i]
	lb.i++

	ip, err := clb.LookupA(lb.serverStr, srv.Target)
	if err != nil {
		return add, err
	}

	return clb.Address{Address: ip, Port: srv.Port}, nil
}
