package letterboxd

import (
	"encoding/json"
	"net/url"

	"github.com/go-resty/resty/v2"
)

// GetAuthToken returns an access token from provided username and password
func (c *Client) GetAuthToken(u string, p string) (AccessToken, error) {
	at := AccessToken{}

	v := url.Values{}
	v.Set("grant_type", "password")
	v.Set("username", u)
	v.Set("password", p)
	s := v.Encode()

	r := c.NewRequest("POST", "/auth/token", s)

	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetHeader("Accept", "application/json").
		Post(r.url)
	if err != nil {
		return at, err
	}

	err = json.Unmarshal(resp.Body(), &at)
	if err != nil {
		return at, err
	}

	return at, nil
}
