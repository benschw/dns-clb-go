package randomclb

import (
	"fmt"
	"github.com/benschw/dns-clb-go/dns"
	"math/rand"
)

func NewRandomClb(lib dns.Lookup) *RandomClb {
	lb := new(RandomClb)
	lb.dnsLib = lib
	return lb
}

type RandomClb struct {
	dnsLib dns.Lookup
}

func (lb *RandomClb) GetAddress(name string) (dns.Address, error) {
	add := dns.Address{}

	srvs, err := lb.dnsLib.LookupSRV(name)
	if err != nil {
		return add, err
	}
	if len(srvs) == 0 {
		return add, fmt.Errorf("no SRV records found")
	}

	//	log.Printf("%+v", srvs)
	srv := srvs[rand.Intn(len(srvs))]

	ip, err := lb.dnsLib.LookupA(srv.Target)
	if err != nil {
		return add, err
	}

	return dns.Address{Address: ip, Port: srv.Port}, nil
}
