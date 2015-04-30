package roundrobinclb

import (
	"fmt"
	"net"
	"sort"

	"github.com/benschw/dns-clb-go/dns"
)

type ByTarget []net.SRV

func (a ByTarget) Len() int           { return len(a) }
func (a ByTarget) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByTarget) Less(i, j int) bool { return a[i].Target < a[j].Target }

func NewRoundRobinClb(lib dns.Lookup) *RoundRobinClb {
	lb := new(RoundRobinClb)
	lb.dnsLib = lib
	lb.i = 0

	return lb
}

type RoundRobinClb struct {
	dnsLib dns.Lookup
	i      int
}

func (lb *RoundRobinClb) GetAddress(name string) (dns.Address, error) {
	add := dns.Address{}

	srvs, err := lb.dnsLib.LookupSRV(name)
	if err != nil {
		return add, err
	}
	sort.Sort(ByTarget(srvs))

	if len(srvs) == 0 {
		return add, fmt.Errorf("no SRV records found")
	}
	if len(srvs)-1 < lb.i {
		lb.i = 0
	}
	//	log.Printf("%d/%d / %+v", lb.i, len(srvs), srvs)
	srv := srvs[lb.i]
	lb.i = lb.i + 1

	ip, err := lb.dnsLib.LookupA(srv.Target)
	if err != nil {
		return add, err
	}

	return dns.Address{Address: ip, Port: srv.Port}, nil
}
