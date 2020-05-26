package nsxt

type NodeCertificateInfo struct {
	// SHA256 of certificate
	CertificateSha256Thumbprint string `json:"certificate_sha256_thumbprint,omitempty"`
	// Certificate content
	Certificate string `json:"certificate,omitempty"`
	// Entity type of this certificate
	EntityType string `json:"entity_type,omitempty"`
}
