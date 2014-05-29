package clb

import (
	"github.com/benschw/consul-clb-go/dns"
)

type LoadBalancer interface {
	GetAddress(name string) (dns.Address, error)
}
