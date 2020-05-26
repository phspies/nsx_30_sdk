package nsxt

// NSX global configs for security purposes, like trust store and trust manager.
type SecurityGlobalConfig struct {
	// Valid Global configuration types
	ResourceType string `json:"resource_type"`
	// When this flag is set to true (for NDcPP compliance) only ca-signed certificates will be allowed to be applied as server certificates.
	CaSignedOnly bool `json:"ca_signed_only,omitempty"`
	// When this flag is set to true, during certificate checking the CRL is fetched and checked whether the certificate is revoked or not.
	CrlCheckingEnabled bool `json:"crl_checking_enabled,omitempty"`
}
