package letterboxd

const (
	// APIBaseURL is used to create a client
	APIBaseURL = "https://api.letterboxd.com/api/v0"
)

// Client provides a common interface and plain configuration for interacting with this library
type Client struct {
	BaseURL   string
	apiKey    string
	apiSecret string
}

// NewClientFromKey creates a letterboxd API client from an auth token
func NewClientFromKey(apiKey string, apiSecret string) *Client {
	return &Client{
		BaseURL:   APIBaseURL,
		apiKey:    apiKey,
		apiSecret: apiSecret,
	}
}

// NewRequest creates a letterboxd API request with all required metadata
// m Method e.g. "GET", "POST"
// e Endpoint e.g. "/auth/token"
// b Body e.g. "{"field": "value"}"
func (c *Client) NewRequest(m string, e string, d string) LbxdReq {
	b := &Builder{}
	r := b.
		Method(m).
		Endpoint(e).
		Data(d).
		APIKey(c.apiKey).
		Build(c)

	return r
}
