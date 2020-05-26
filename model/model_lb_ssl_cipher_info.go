package nsxt

// SSL cipher
type LbSslCipherInfo struct {
	// Default SSL cipher flag
	IsDefault bool `json:"is_default"`
	// Secure/insecure SSL cipher flag
	IsSecure bool `json:"is_secure"`
	// Several cipher groups might contain the same cipher suite, each cipher suite could have multiple cipher group labels. 
	CipherGroupLabels []string `json:"cipher_group_labels,omitempty"`
	// SSL cipher
	Cipher string `json:"cipher"`
}
