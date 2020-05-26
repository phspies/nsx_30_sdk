package nsxt

// Base type for various login credential types
type LoginCredential struct {
	// Possible values are 'UsernamePasswordLoginCredential', 'VerifiableAsymmetricLoginCredential', 'SessionLoginCredential'.
	CredentialType string `json:"credential_type"`
}
