package clb

import (
	"fmt"
	"github.com/miekg/dns"
	"net"
	"strconv"
	"strings"
)

type LoadBalancer interface {
	GetAddress(name string) (Address, error)
}

type DNSServerConfig struct {
	Address string
	Port    string
}

type Address struct {
	Address string
	Port    uint16
}

func LookupSRV(cfg DNSServerConfig, name string) ([]net.SRV, error) {
	var srvs = make([]net.SRV, 0, 10)
	answer, err := lookupType(cfg, name, "SRV")
	if err != nil {
		return srvs, err
	}
	for _, v := range answer.Answer {
		//log.Printf("%+v", v.Header())
		parts := strings.Split(v.String(), "\t")
		info := parts[4]
		infoParts := strings.Split(info, " ")

		n := len(srvs)
		srvs = srvs[0 : n+1]
		port, err := strconv.ParseUint(infoParts[2], 0, 16)
		if err != nil {
			return srvs, err
		}
		srvs[n] = net.SRV{Target: infoParts[3], Port: uint16(port)}

	}
	return srvs, nil
}
func LookupA(cfg DNSServerConfig, name string) (string, error) {
	answer, err := lookupType(cfg, name, "A")
	if err != nil {
		return "", err
	}
	parts := strings.Split(answer.Answer[0].String(), "\t")
	return parts[len(parts)-1], nil
}

func lookupType(server DNSServerConfig, name string, recordType string) (*dns.Msg, error) {
	qType, ok := dns.StringToType[recordType]
	if !ok {
		return nil, fmt.Errorf("Invalid type '%s'", recordType)
	}
	name = dns.Fqdn(name)

	client := &dns.Client{}
	msg := &dns.Msg{}
	msg.SetQuestion(name, qType)

	serverStr := fmt.Sprintf("%s:%s", server.Address, server.Port)
	response, err := lookup(msg, client, serverStr, false)
	if err != nil {
		return nil, fmt.Errorf("Couldn't resolve %s: No server responded", name)
	}
	return response, nil

}

func lookup(msg *dns.Msg, client *dns.Client, server string, edns bool) (*dns.Msg, error) {
	if edns {
		opt := &dns.OPT{
			Hdr: dns.RR_Header{
				Name:   ".",
				Rrtype: dns.TypeOPT,
			},
		}
		opt.SetUDPSize(dns.DefaultMsgSize)
		msg.Extra = append(msg.Extra, opt)
	}

	response, _, err := client.Exchange(msg, server)

	if err != nil {
		return nil, err
	}

	if msg.Id != response.Id {
		return nil, fmt.Errorf("DNS ID mismatch, request: %d, response: %d", msg.Id, response.Id)
	}

	return response, nil
}
