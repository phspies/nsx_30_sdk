package nsxt

// A login credential specifying a username and password
type UsernamePasswordLoginCredential struct {
	// Possible values are 'UsernamePasswordLoginCredential', 'VerifiableAsymmetricLoginCredential', 'SessionLoginCredential'.
	CredentialType string `json:"credential_type"`
	// The username for login
	Username string `json:"username,omitempty"`
	// The authentication password for login
	Password string `json:"password,omitempty"`
	// Thumbprint of the login server
	Thumbprint string `json:"thumbprint,omitempty"`
}
