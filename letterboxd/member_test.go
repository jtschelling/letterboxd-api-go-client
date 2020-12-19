package letterboxd

import (
	"testing"
)

func TestGetSelf(t *testing.T) {
	c := NewClientFromKey("FAKE_KEY", "FAKE_SECRET")

	at, _ := c.GetUserAccessToken("USERNAME", "PASSWORD")
	r, err := c.GetSelf(at)

	if err != nil {
		t.Error(r)
	}
}
