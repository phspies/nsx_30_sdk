package nsxt

// SSL protocol
type LbSslProtocolInfo struct {
	// Default SSL protocol flag
	IsDefault bool `json:"is_default"`
	// Secure/insecure SSL protocol flag
	IsSecure bool `json:"is_secure"`
	// SSL protocol
	Protocol string `json:"protocol"`
}
