package letterboxd

import (
	"testing"
)

func TestNewClientFromKey(t *testing.T) {
	url := "https://api.letterboxd.com/api/v0"
	k := "FAKE_KEY"
	sk := "FAKE_SECRET"
	c := NewClientFromKey(k, sk)
	if c.apiKey != k {
		t.Errorf("Client returned wrong apiKey field\n\twanted %s\n\tgot %s", k, c.apiKey)
	}
	if c.apiSecret != sk {
		t.Errorf("Client returned wrong apiSecret field\n\twanted %s\n\tgot %s", sk, c.apiSecret)
	}
	if c.BaseURL != url {
		t.Errorf("Client returned wrong BaseURL field\n\twanted %s\n\tgot %s", url, c.BaseURL)
	}
}
