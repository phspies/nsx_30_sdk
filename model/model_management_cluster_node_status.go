package nsxt

type ManagementClusterNodeStatus struct {
	// Status of this node's connection to the management cluster
	MgmtClusterStatus string `json:"mgmt_cluster_status,omitempty"`
}
