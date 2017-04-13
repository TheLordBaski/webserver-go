package server

import (
	"fmt"
	"testing"
)

// TestHttpAddress ensures the correct address is returned.
func TestHttpAddress(t *testing.T) {
	i := Info{
		Hostname: "example.com",
		HTTPPort: 8080,
	}

	expected := fmt.Sprintf("%v:%d", i.Hostname, i.HTTPPort)
	received := httpAddress(i)

	if expected != received {
		t.Errorf("\n got: %v\nwant: %v", received, expected)
	}
}

