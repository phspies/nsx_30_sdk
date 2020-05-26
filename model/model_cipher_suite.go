package nsxt

// TLS cipher suite
type CipherSuite struct {
	// Enable status for this cipher suite
	Enabled bool `json:"enabled"`
	// Name of the TLS cipher suite
	Name string `json:"name"`
}
