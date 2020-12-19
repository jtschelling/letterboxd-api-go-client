package letterboxd

import (
	"testing"
)

func TestGetUserAccessToken(t *testing.T) {
	c := NewClientFromKey("FAKE_KEY", "FAKE_SECRET")

	r, err := c.GetUserAccessToken("USERNAME", "PASSWORD")

	if err != nil {
		t.Error(r)
	}
}
