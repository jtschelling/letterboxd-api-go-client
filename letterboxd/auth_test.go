package letterboxd

import (
	"testing"
)

// TODO: This isn't testing correctly. Needs a mocked endpoint
func TestGetAuthToken(t *testing.T) {
	c := NewClientFromKey("FAKE_KEY", "FAKE_SECRET")
	at, err := c.GetAuthToken("USERNAME", "PASSWORD")

	if err != nil {
		t.Error(at)
	}
}
