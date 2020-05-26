package nsxt

type CertificateKeyPair struct {
	// The private key must include the enclosing \"-----BEGIN RSA PRIVATE KEY-----\" and \"-----END RSA PRIVATE KEY-----\". An empty string is returned in read responses.
	RsaPrivateKey string `json:"rsa_private_key,omitempty"`
	Certificate *SecurityCertificate `json:"certificate"`
}
