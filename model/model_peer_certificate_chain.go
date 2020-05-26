package nsxt

// The certificate chain presented by a remote TLS service.
type PeerCertificateChain struct {
	// List of X509Certificates.
	Details []X509Certificate `json:"details,omitempty"`
	// PEM encoded certificate data.
	PemEncoded string `json:"pem_encoded"`
}
