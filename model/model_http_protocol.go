package nsxt

type HttpProtocol struct {
	// Protocol name
	Name string `json:"name"`
	AuthenticationScheme *BasicAuthenticationScheme `json:"authentication_scheme,omitempty"`
}
