package clb

import (
	"fmt"
	"log"
	"testing"
)

var _ = fmt.Print // For debugging; delete when done.
var _ = log.Print // For debugging; delete when done.

var tmpPath = "/tmp/pid-tests"
var testSvc = "ServiceName"
var testId = "jhgsd765asd"

func Test_LookupConsul_ReturnsConsulAddress(t *testing.T) {
	// given
	port := "8300"

	// when
	srvName := "consul.service.consul"
	address, err := LookupAddress(srvName)

	// then
	if err != nil {
		t.Error(err)
	}
	if port != address.Port {
		t.Errorf("port %s not expected", address.Port)
	}
	fmt.Printf("%+v", address)
}
