package ttlcache

import (
	"github.com/benschw/dns-clb-go/dns"
	"net"
	"time"
)

func NewTtlCache(lib dns.Lookup, ttl int) *TtlCache {
	c := new(TtlCache)
	c.lib = lib
	c.ttl = ttl
	c.lastUpdate = 0

	return c
}

type TtlCache struct {
	lib        dns.Lookup
	ttl        int
	lastUpdate int32
	srvs       []net.SRV
	as         map[string]string
}

func (l *TtlCache) LookupSRV(name string) ([]net.SRV, error) {
	err := l.checkCache()
	if err != nil {
		return nil, err
	}

	if len(l.srvs) == 0 {
		l.srvs, err = l.lib.LookupSRV(name)
		if err != nil {
			return nil, err
		}
	}
	return l.srvs, nil
}

func (l *TtlCache) LookupA(name string) (string, error) {
	err := l.checkCache()
	if err != nil {
		return "", err
	}

	_, ok := l.as[name]
	if !ok {
		l.as[name], err = l.lib.LookupA(name)
		if err != nil {
			return "", err
		}
	}

	return l.as[name], nil
}

func (l *TtlCache) checkCache() error {
	now := int32(time.Now().Unix())
	if l.lastUpdate+int32(l.ttl) < now {
		l.lastUpdate = now
		l.srvs = []net.SRV{}
		l.as = map[string]string{}
	}
	return nil
}
