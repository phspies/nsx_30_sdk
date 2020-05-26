package nsxt

// The basic node info of management plane node
type ManagementPlaneBaseNodeInfo struct {
	// Management plane node UUID
	Uuid string `json:"uuid,omitempty"`
	// The IP address of MP node
	MgmtClusterListenIpAddress string `json:"mgmt_cluster_listen_ip_address,omitempty"`
}
