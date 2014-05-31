package randomclb

import (
	"fmt"
	"github.com/benschw/dns-clb-go/dns"
	"log"
	"testing"
)

var _ = fmt.Print // For debugging; delete when done.
var _ = log.Print // For debugging; delete when done.

//strconv.FormatInt(int64(srv.Port), 10)
func TestRandomLookup(t *testing.T) {
	// given
	srvName := "foo.service.fliglio.com"
	lib := dns.NewLookupLib("8.8.8.8:53")
	c := NewRandomClb(lib)

	// when
	address, err := c.GetAddress(srvName)

	// then
	if err != nil {
		t.Error(err)
	}

	if address.Port == 8001 && address.Address == "0.1.2.3" {
		return
	} else if address.Port == 8002 && address.Address == "4.5.6.7" {
		return
	} else {
		t.Errorf("port '%d' not expected with address: '%s'", address.Port, address.Address)
	}
}
