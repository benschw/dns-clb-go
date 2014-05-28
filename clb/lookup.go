package clb

import (
	"fmt"
	"github.com/miekg/dns"
	"math/rand"
	"strings"
)

const resolvConf = "/etc/resolv.conf"

type DNSServer struct {
	Address string
	Port    string
}

type Address struct {
	Address string
	Port    string
}

func LookupAddress(server DNSServer, name string) (Address, error) {
	add := Address{}

	srvs, err := LookupSRV(server, name)
	if err != nil {
		return add, err
	}
	//	log.Printf("%+v", srvs)

	srv := srvs[rand.Intn(len(srvs))]

	aAdd, err := LookupA(server, srv.Address)
	if err != nil {
		return add, err
	}

	return Address{Address: aAdd.Address, Port: srv.Port}, nil
}
func LookupSRV(server DNSServer, name string) ([]Address, error) {
	var srvs = make([]Address, 0, 10)
	answer, err := Lookup(server, name, "SRV")
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
		srvs[n] = Address{Address: infoParts[3], Port: infoParts[2]}

	}
	return srvs, nil
}
func LookupA(server DNSServer, name string) (Address, error) {
	add := Address{}
	answer, err := Lookup(server, name, "A")
	if err != nil {
		return add, err
	}
	parts := strings.Split(answer.Answer[0].String(), "\t")
	add.Address = parts[len(parts)-1]

	return add, nil
}

func Lookup(server DNSServer, name string, recordType string) (*dns.Msg, error) {
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
