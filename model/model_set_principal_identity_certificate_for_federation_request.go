package nsxt

// Data for setting a principal identity certificate
type SetPrincipalIdentityCertificateForFederationRequest struct {
	// Service type for which the certificate should be used.
	ServiceType string `json:"service_type,omitempty"`
	// Id of the certificate
	CertId string `json:"cert_id,omitempty"`
}
