package nsxt

type InitiateClusterRestoreRequest struct {
	// Timestamp of the backed-up configuration from which the appliance will be restored 
	Timestamp int64 `json:"timestamp,omitempty"`
	// Unique id of the backed-up configuration from which the appliance will be restored 
	NodeId string `json:"node_id,omitempty"`
	// IP address or FQDN of the node from which the backup was taken
	IpAddress string `json:"ip_address,omitempty"`
}
