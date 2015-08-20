package clb

import (
	"fmt"

	"github.com/benschw/dns-clb-go/dns"
	"github.com/benschw/dns-clb-go/randomclb"
	"github.com/benschw/dns-clb-go/roundrobinclb"
	"github.com/benschw/dns-clb-go/ttlcache"
)

type LoadBalancerType int

const (
	Random     LoadBalancerType = iota
	RoundRobin LoadBalancerType = iota
)

type CacheType int

const (
	None CacheType = iota
	Ttl  CacheType = iota
)

type LoadBalancer interface {
	GetAddress(name string) (dns.Address, error)
}

func New() LoadBalancer {
	return NewDefaultClb(RoundRobin)
}

func NewDefaultClb(lbType LoadBalancerType) LoadBalancer {
	lib := dns.NewDefaultLookupLib()

	return buildClb(lib, lbType)
}

func NewClb(address string, port string, lbType LoadBalancerType) LoadBalancer {
	lib := dns.NewLookupLib(fmt.Sprintf("%s:%s", address, port))

	return buildClb(lib, lbType)
}

func NewTtlCacheClb(address string, port string, lbType LoadBalancerType, ttl int) LoadBalancer {
	lib := dns.NewLookupLib(fmt.Sprintf("%s:%s", address, port))
	cache := ttlcache.NewTtlCache(lib, ttl)

	return buildClb(cache, lbType)
}

func buildClb(lib dns.Lookup, lbType LoadBalancerType) LoadBalancer {
	switch lbType {
	case RoundRobin:
		return NewRoundRobinClb(lib)
	case Random:
		return NewRandomClb(lib)
	}
	return nil
}

func NewRoundRobinClb(lib dns.Lookup) *roundrobinclb.RoundRobinClb {
	return roundrobinclb.NewRoundRobinClb(lib)
}

func NewRandomClb(lib dns.Lookup) *randomclb.RandomClb {
	return randomclb.NewRandomClb(lib)
}
