package nsxt

// Configuration where backup files are stored for restore
type RestoreConfiguration struct {
	RemoteFileServer *RemoteFileServer `json:"remote_file_server"`
	// Passphrase used to encrypt backup files.
	Passphrase string `json:"passphrase,omitempty"`
}
