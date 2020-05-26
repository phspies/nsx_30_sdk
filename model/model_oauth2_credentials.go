package nsxt

// Oauth2 Account Credentials
type Oauth2Credentials struct {
	// Client secret, that will be used for authentication in AWS environment. Can be some passphrase.
	ClientSecret string `json:"client_secret,omitempty"`
	// Client ID, that will be used for authentication in AWS environment,
	ClientId string `json:"client_id"`
}
