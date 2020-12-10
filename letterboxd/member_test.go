package letterboxd

import (
	"testing"
)

func TestGetSelf(t *testing.T) {
	c := NewClientFromKey("FAKE_KEY", "FAKE_SECRET")

	at, _ := c.getUserAccessToken("USERNAME", "PASSWORD")
	r, err := c.getSelf(at)

	if err != nil {
		t.Error(r)
	}
}
