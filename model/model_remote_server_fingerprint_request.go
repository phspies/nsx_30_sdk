package nsxt

// Remote server
type RemoteServerFingerprintRequest struct {
	// Server port
	Port int64 `json:"port,omitempty"`
	// Remote server hostname or IP address
	Server string `json:"server"`
}
