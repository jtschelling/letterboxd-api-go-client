package letterboxd

// AccessToken holds information to sign every request with a valid access token
// data provided by the getAccessToken() function
type AccessToken struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}
