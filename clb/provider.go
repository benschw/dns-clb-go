package clb

import "github.com/benschw/dns-clb-go/dns"

type AddressProvider interface {
	GetAddress() (dns.Address, error)
}

func NewAddressProvider(address string) *SRVAddressProvider {
	return &SRVAddressProvider{
		Lb:      New(),
		Address: address,
	}
}

type SRVAddressProvider struct {
	Lb      LoadBalancer
	Address string
}

func (s *SRVAddressProvider) GetAddress() (dns.Address, error) {
	return s.Lb.GetAddress(s.Address)
}
