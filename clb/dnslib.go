package clb

import (
	"fmt"
	"github.com/miekg/dns"
)

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
