package nsxt

// Cluster backup details
type ClusterBackupInfo struct {
	// timestamp of the cluster backup file
	Timestamp int64 `json:"timestamp,omitempty"`
	// ID of the node from which the backup was taken
	NodeId string `json:"node_id,omitempty"`
	// IP address or FQDN of the node from which the backup was taken
	IpAddress string `json:"ip_address,omitempty"`
	// Type of restore allowed
	RestoreType []string `json:"restore_type,omitempty"`
}
