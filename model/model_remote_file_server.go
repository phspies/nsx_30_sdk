package nsxt

// Remote file server
type RemoteFileServer struct {
	// Remote server directory to copy bundle files to
	DirectoryPath string `json:"directory_path"`
	Protocol *FileTransferProtocol `json:"protocol"`
	// Server port
	Port int64 `json:"port,omitempty"`
	// Remote server hostname or IP address
	Server string `json:"server"`
}
