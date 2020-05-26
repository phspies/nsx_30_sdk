package nsxt

// Result of checking a certificate
type CertificateCheckingStatus struct {
	// Status of the checked certificate.
	Status string `json:"status,omitempty"`
	// Error message when checking the certificate.
	ErrorMessage string `json:"error_message,omitempty"`
}
