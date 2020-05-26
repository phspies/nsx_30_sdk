package nsxt

// Remote file server
type SupportBundleRemoteFileServer struct {
	// Uploads to the remote file server performed by the manager
	ManagerUploadOnly bool `json:"manager_upload_only,omitempty"`
	// Remote server directory to copy bundle files to
	DirectoryPath string `json:"directory_path"`
	Protocol *SupportBundleFileTransferProtocol `json:"protocol"`
	// Server port
	Port int64 `json:"port,omitempty"`
	// Remote server hostname or IP address
	Server string `json:"server"`
}
