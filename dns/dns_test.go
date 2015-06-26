package dns

import (
	"net"
	"testing"
)

func TestLookupShoudFailWithBadNS(t *testing.T) {
	lib := NewLookupLib("foo:9999")

	_, err := lib.LookupA("foo")

	if err == nil {
		t.Error("looking up foo on foo:9999 should product an error")
	}
}

func TestLookupShoudFailWithBadHost(t *testing.T) {
	lib := NewLookupLib("8.8.8.8:53")

	_, err := lib.LookupA("foo")

	if err == nil {
		t.Error("looking up foo on 8.8.8.8:53 should product an error")
	}
}

func TestLookupShouldResolveARecord(t *testing.T) {
	lib := NewLookupLib("8.8.8.8:53")

	address, err := lib.LookupA("github.com")

	if err != nil {
		t.Error(err)
	}

	ip := net.ParseIP(address)
	if ip.To4() == nil {
		t.Errorf("address '%s' not valid", address)
	}

}
