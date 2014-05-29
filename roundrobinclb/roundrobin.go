package roundrobinclb

import (
	"fmt"
	"github.com/benschw/consul-clb-go/dns"
)

func NewRoundRobinClb(address string, port string) *RoundRobinClb {
	lb := new(RoundRobinClb)
	lb.serverStr = fmt.Sprintf("%s:%s", address, port)
	lb.i = 0

	return lb
}

type RoundRobinClb struct {
	serverStr string
	i         int
}

func (lb *RoundRobinClb) GetAddress(name string) (dns.Address, error) {
	add := dns.Address{}

	srvs, err := dns.LookupSRV(lb.serverStr, name)
	if err != nil {
		return add, err
	}
	if len(srvs) == 0 {
		return add, fmt.Errorf("no SRV records found")
	}
	//	log.Printf("%+v", srvs)
	if len(srvs)-1 > lb.i {
		lb.i = 0
	}
	srv := srvs[lb.i]
	lb.i = lb.i + 1

	ip, err := dns.LookupA(lb.serverStr, srv.Target)
	if err != nil {
		return add, err
	}

	return dns.Address{Address: ip, Port: srv.Port}, nil
}
