package nsxt

// Remote server
type RemoteServerFingerprint struct {
	// SSH fingerprint of server
	SshFingerprint string `json:"ssh_fingerprint"`
	// Server port
	Port int64 `json:"port,omitempty"`
	// Remote server hostname or IP address
	Server string `json:"server"`
}
