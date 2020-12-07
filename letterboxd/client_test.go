package letterboxd

import (
	"testing"
)

func TestNewClientFromKey(t *testing.T) {
	k := "FAKE_KEY"
	sk := "FAKE_SECRET"
	c := NewClientFromKey(k, sk)
	if c.apiKey != k {
		t.Errorf("Client returned wrong apiKey field\n\twanted %s\n\tgot %s", k, c.apiKey)
	}
	if c.apiSecret != sk {
		t.Errorf("Client returned wrong apiSecret field\n\twanted %s\n\tgot %s", sk, c.apiSecret)
	}
}

func TestNewRequest(t *testing.T) {
	k := "FAKE_KEY"
	sk := "FAKE_SECRET"
	c := NewClientFromKey(k, sk)
	r := c.NewRequest("GET", "/fake", "")

	if r.method != "GET" {
		t.Errorf("Request container wrong method\n\twanted %s\n\tgo %s", "GET", r.method)
	}
}
