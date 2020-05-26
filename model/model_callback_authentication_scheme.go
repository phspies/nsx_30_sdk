package nsxt

// CallbackAuthenticationScheme describes how notification requests should authenticate to the server.
type CallbackAuthenticationScheme struct {
	// Username to use if scheme_name is BASIC_AUTH.
	Username string `json:"username,omitempty"`
	// Certificate ID with a valid certificate and private key, procured from trust-management API.
	CertificateId string `json:"certificate_id,omitempty"`
	// Authentication scheme to use when making notification requests to the partner console. Specify one of BASIC_AUTH or CERTIFICATE.
	SchemeName string `json:"scheme_name"`
	// Password to use if scheme_name is BASIC_AUTH.
	Password string `json:"password,omitempty"`
}
