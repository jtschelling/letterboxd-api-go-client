package letterboxd

import (
	"testing"
)

func TestNewClientFromKey(t *testing.T) {
	k := "FAKE_KEY"
	c := NewClientFromKey(k, "FAKE_SECRET")
	if c.apiKey != k {
		t.Errorf("Client returned wrong apiKey field\n\twanted %s\n\tgot %s", k, c.apiKey)
	}
}
