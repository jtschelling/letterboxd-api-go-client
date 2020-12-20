package letterboxd

import (
	"encoding/json"

	"github.com/go-resty/resty/v2"
)

// GetMe retrieves information from the /me endpoint
func (c *Client) GetMe(at string) (MemberAccount, error) {
	r := c.NewRequest("GET", "/me", "")

	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetHeader("Accept", "application/json").
		SetAuthToken(at).
		Get(r.url)
	if err != nil {
		return MemberAccount{}, err
	}

	ma := MemberAccount{}
	err = json.Unmarshal(resp.Body(), &ma)
	if err != nil {
		return MemberAccount{}, err
	}

	return ma, nil
}
