package nsxt

type CopyRemoteFileProperties struct {
	// URI of file to copy
	Uri string `json:"uri"`
	// Server port
	Port int64 `json:"port,omitempty"`
	// Remote server hostname or IP address
	Server string `json:"server"`
}
