package clb

import (
	"fmt"
	"github.com/miekg/dns"
	"math/rand"
	"strings"
)

const resolvConf = "/etc/resolv.conf"

type Address struct {
	Address string
	Port    string
}

func LookupAddress(name string) (Address, error) {
	add := Address{}
	t, ok := dns.StringToType["SRV"]
	if !ok {
		return add, fmt.Errorf("Invalid type 'SRV'")
	}

	answer, err := Lookup(t, name)
	if err != nil {
		return add, err
	}
	var srvs = make([]Address, 0, 10)
	for _, v := range answer.Answer {
		//log.Printf("%+v", v.Header())
		parts := strings.Split(v.String(), "\t")
		info := parts[4]
		infoParts := strings.Split(info, " ")

		n := len(srvs)
		srvs = srvs[0 : n+1]
		srvs[n] = Address{Address: infoParts[3], Port: infoParts[2]}

	}
	//	log.Printf("%+v", srvs)

	srv := srvs[rand.Intn(len(srvs))]

	ip, err := LookupA(srv.Address)
	if err != nil {
		return add, err
	}

	return Address{Address: ip, Port: srv.Port}, nil
}
func LookupA(name string) (string, error) {
	t, ok := dns.StringToType["A"]
	if !ok {
		return "", fmt.Errorf("Invalid type 'A'")
	}

	answer, err := Lookup(t, name)
	if err != nil {
		return "", err
	}
	parts := strings.Split(answer.Answer[0].String(), "\t")
	ip := parts[len(parts)-1]
	//	log.Printf("%+v", answer.Answer)
	return ip, nil
}

func Lookup(qType uint16, name string) (*dns.Msg, error) {
	name = dns.Fqdn(name)

	client := &dns.Client{}
	msg := &dns.Msg{}
	msg.SetQuestion(name, qType)

	server := fmt.Sprintf("%s:%s", "127.0.0.1", "8600")
	response, err := lookup(msg, client, server, false)
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
