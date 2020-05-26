package nsxt

type VerifiableAsymmetricLoginCredential struct {
	// Possible values are 'UsernamePasswordLoginCredential', 'VerifiableAsymmetricLoginCredential', 'SessionLoginCredential'.
	CredentialType string `json:"credential_type"`
	// Asymmetric login credential
	AsymmetricCredential string `json:"asymmetric_credential,omitempty"`
	// Credential verifier
	CredentialVerifier string `json:"credential_verifier,omitempty"`
	// Credential key
	CredentialKey string `json:"credential_key,omitempty"`
}
