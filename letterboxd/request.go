package letterboxd

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/google/uuid"
)

// LbxdReq defines a single API request
type LbxdReq struct {
	method    string
	endpoint  string
	body      string
	apiKey    string
	nonce     string
	timestamp string
	signature string
	url       string
	data      string
}

// Builder helps facilitate a new request object
type Builder struct {
	r LbxdReq
}

// Build creates all non-user input fields, generates a request signature, and returns the completed request object
func (b *Builder) Build(c *Client) LbxdReq {
	b.nonce()
	b.timestamp()
	b.buildSignature(c)
	b.buildURL(c)
	return b.r
}

// Method sets the request method
func (b *Builder) Method(m string) *Builder {
	b.r.method = m
	return b
}

// Endpoint sets the request endpoint
func (b *Builder) Endpoint(e string) *Builder {
	b.r.endpoint = e
	return b
}

// Body sets the body struct field for requests that require it
func (b *Builder) Body(d string) *Builder {
	b.r.body = d
	return b
}

// APIKey sets the request api key
func (b *Builder) APIKey(a string) *Builder {
	b.r.apiKey = a
	return b
}

// Data sets the request data field, added to the URL
func (b *Builder) Data(d string) *Builder {
	b.r.data = d
	return b
}

// timestamp create the request timestamp in Unix seconds
func (b *Builder) timestamp() *Builder {
	b.r.timestamp = createTimestamp()
	return b
}

// nonce create the request UUID
func (b *Builder) nonce() *Builder {
	b.r.nonce = createNonce()
	return b
}

// buildURL forms the final URL with base, endpoint, and signature information
func (b *Builder) buildURL(c *Client) *Builder {
	var sb strings.Builder

	sb.WriteString(c.BaseURL)
	sb.WriteString(b.r.endpoint)
	sb.WriteString(b.addURLMetadata(c))
	sb.WriteString("&")
	sb.WriteString(b.r.data)
	sb.WriteString("&signature=")
	sb.WriteString(b.r.signature)
	b.r.url = sb.String()
	return b
}

// generateReqSignature computes the signature hash required by the Letterboxd API for every request
// http://api-docs.letterboxd.com/#signing
func (b *Builder) buildSignature(c *Client) *Builder {
	saltStr := b.buildSalt(c)

	h := hmac.New(sha256.New, []byte(c.apiSecret))
	_, err := h.Write([]byte(saltStr))
	if err != nil {
		fmt.Println("Error in buildSignature")
	}

	hexResult := hex.EncodeToString(h.Sum(nil))

	s := strings.ToLower(hexResult)
	b.r.signature = s
	return b
}

// builtSalt computes a Letterboxd specified salt string for use in generating the request signature
// http://api-docs.letterboxd.com/#signing
func (b *Builder) buildSalt(c *Client) string {
	var sb strings.Builder

	sb.WriteString(b.r.method)
	sb.WriteString("\u0000")
	sb.WriteString(c.BaseURL)
	sb.WriteString(b.r.endpoint)

	// add apikey nonce and timestamp as endpoint parameters
	sb.WriteString(b.addURLMetadata(c))
	sb.WriteString("&")
	sb.WriteString(b.r.data)

	sb.WriteString("\u0000")
	sb.WriteString(b.r.body)

	// fmt.Printf("\n\n%+q\n\n", sb.String())

	return sb.String()
}

func (b *Builder) addURLMetadata(c *Client) string {
	v := url.Values{}
	v.Set("apikey", c.apiKey)
	v.Set("nonce", b.r.nonce)
	v.Set("timestamp", b.r.timestamp)
	e := v.Encode()

	return "?" + e
}

// createNonce generates a UUID for a request
func createNonce() string {
	return uuid.New().String()
}

// createTimestamp generates a timestamp for a request
func createTimestamp() string {
	return fmt.Sprint(time.Now().Unix())
}
