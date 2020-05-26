package nsxt

// Details of session based login credential to login to server.
type SessionLoginCredential struct {
	// Possible values are 'UsernamePasswordLoginCredential', 'VerifiableAsymmetricLoginCredential', 'SessionLoginCredential'.
	CredentialType string `json:"credential_type"`
	// The session_id to login to server.
	SessionId string `json:"session_id,omitempty"`
	// Thumbprint of the login server.
	Thumbprint string `json:"thumbprint,omitempty"`
}
