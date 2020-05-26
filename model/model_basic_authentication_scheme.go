package nsxt

type BasicAuthenticationScheme struct {
	// Authentication scheme name
	SchemeName string `json:"scheme_name"`
	// User name to authenticate with
	Username string `json:"username"`
	// Password to authenticate with
	Password string `json:"password,omitempty"`
}
