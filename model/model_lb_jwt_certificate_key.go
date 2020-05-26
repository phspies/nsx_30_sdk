package nsxt

// The key is used to specify certificate which is used to verify the signature of JWT tokens. 
type LbJwtCertificateKey struct {
	// The property is used to identify JWT key type. 
	Type_ string `json:"type"`
	// Certificate identifier
	CertificateId string `json:"certificate_id"`
}
