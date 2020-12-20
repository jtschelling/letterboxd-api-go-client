package letterboxd

import "testing"

// TODO: This isn't testing correctly. Needs a mocked endpoint
func TestGetMe(t *testing.T) {
	c := NewClientFromKey("FAKE_KEY", "FAKE_SECRET")

	at, _ := c.GetAuthToken("USERNAME", "PASSWORD")
	r, err := c.GetMe(at.AccessToken)

	if err != nil {
		t.Error(r)
	}
}
