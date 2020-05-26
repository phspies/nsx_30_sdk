package nsxt

type HttpsProtocol struct {
	// Protocol name
	Name string `json:"name"`
	// SSL thumbprint of server
	Sha256Thumbprint string `json:"sha256_thumbprint"`
	AuthenticationScheme *BasicAuthenticationScheme `json:"authentication_scheme,omitempty"`
}
