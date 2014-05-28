package clb

import (
	"fmt"
	"github.com/miekg/dns"
	"net"
)

type LoadBalancer interface {
	GetAddress(name string) (Address, error)
}

type Address struct {
	Address string
	Port    uint16
}

func (a Address) String() string {
	return fmt.Sprintf("%s:%d", a.Address, a.Port)
}

func LookupSRV(serverString string, name string) ([]net.SRV, error) {
	var srvs = make([]net.SRV, 0)
	answer, err := lookupType(serverString, name, "SRV")
	if err != nil {
		return srvs, err
	}
	return parseSRVAnswer(answer)
}

func LookupA(serverString string, name string) (string, error) {
	answer, err := lookupType(serverString, name, "A")
	if err != nil {
		return "", err
	}
	return parseAAnswer(answer)
}

func parseSRVAnswer(answer *dns.Msg) ([]net.SRV, error) {
	var srvs = make([]net.SRV, 0)
	for _, v := range answer.Answer {
		if srv, ok := v.(*dns.SRV); ok {
			srvs = append(srvs, net.SRV{
				Priority: srv.Priority,
				Weight:   srv.Weight,
				Port:     srv.Port,
				Target:   srv.Target,
			})
		}
	}
	return srvs, nil
}

func parseAAnswer(answer *dns.Msg) (string, error) {
	if a, ok := answer.Answer[0].(*dns.A); ok {
		return string(a.A), nil
	}
	return "", fmt.Errorf("Could not parse A record")
}
