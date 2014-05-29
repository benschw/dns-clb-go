package dns

import (
	"fmt"
	"github.com/miekg/dns"
	"net"
)

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

		return a.A.String(), nil

		//		return string(a.A[:n]), nil
	}
	return "", fmt.Errorf("Could not parse A record")
}

func lookupType(serverStr string, name string, recordType string) (*dns.Msg, error) {
	// try a connection with a udp connection first
	return lookup(serverStr, name, recordType, "")
}

func lookup(serverStr string, name string, recordType string, connType string) (*dns.Msg, error) {
	qType, ok := dns.StringToType[recordType]
	if !ok {
		return nil, fmt.Errorf("Invalid type '%s'", recordType)
	}
	name = dns.Fqdn(name)

	client := &dns.Client{Net: connType}

	msg := &dns.Msg{}
	msg.SetQuestion(name, qType)

	response, _, err := client.Exchange(msg, serverStr)

	if err != nil {
		// retry lookup with a tcp connection
		return lookup(serverStr, name, recordType, "tcp")
	}

	if msg.Id != response.Id {
		return nil, fmt.Errorf("DNS ID mismatch, request: %d, response: %d", msg.Id, response.Id)
	}

	return response, nil
}
