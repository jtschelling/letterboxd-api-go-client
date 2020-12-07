package letterboxd

import (
	"fmt"
	"testing"
)

func TestGetUserAccessToken(t *testing.T) {
	c := NewClientFromKey("FAKE_KEY", "FAKE_SECRET")

	r, err := c.getUserAccessToken("USERNAME", "PASSWORD")

	fmt.Println(r)

	if err != nil {
		t.Error(r)
	}
}
