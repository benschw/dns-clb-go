package randomclb

import (
	"fmt"
	"github.com/benschw/consul-clb-go/dns"
	"math/rand"
)

func NewRandomClb(address string, port string) *RandomClb {
	lb := new(RandomClb)
	lb.serverStr = fmt.Sprintf("%s:%s", address, port)

	return lb
}

type RandomClb struct {
	serverStr string
}

func (lb *RandomClb) GetAddress(name string) (dns.Address, error) {
	add := dns.Address{}

	srvs, err := dns.LookupSRV(lb.serverStr, name)
	if err != nil {
		return add, err
	}
	//	log.Printf("%+v", srvs)

	srv := srvs[rand.Intn(len(srvs))]

	ip, err := dns.LookupA(lb.serverStr, srv.Target)
	if err != nil {
		return add, err
	}

	return dns.Address{Address: ip, Port: srv.Port}, nil
}
