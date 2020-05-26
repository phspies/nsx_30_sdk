package nsxt

// TLS protocol version
type ProtocolVersion struct {
	// Enable status for this protocol version
	Enabled bool `json:"enabled"`
	// Name of the TLS protocol version
	Name string `json:"name"`
}
