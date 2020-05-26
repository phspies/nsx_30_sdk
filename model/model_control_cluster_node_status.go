package nsxt

type ControlClusterNodeStatus struct {
	MgmtConnectionStatus *MgmtConnStatus `json:"mgmt_connection_status,omitempty"`
	// Status of this node's connection to the control cluster
	ControlClusterStatus string `json:"control_cluster_status,omitempty"`
}
