package letterboxd

import (
	"encoding/json"
	"net/url"

	"github.com/go-resty/resty/v2"
)

// GetUserAccessToken returns an access token from provided username and password
func (c *Client) GetUserAccessToken(u string, p string) (string, error) {
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
		return "", err
	}

	at := AccessToken{}
	err = json.Unmarshal(resp.Body(), &at)
	if err != nil {
		return "", err
	}

	return at.AccessToken, nil
}
