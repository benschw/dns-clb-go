package dns

import (
	"testing"
)

func TestResolveShouldFail(t *testing.T) {
	lib := NewLookupLib("foo:9999")

	_, err := lib.LookupA("foo")

	if err == nil {
		t.Error("looking up foo on foo:9999 should product an error")
	}
}
