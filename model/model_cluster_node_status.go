package nsxt

type ClusterNodeStatus struct {
	SystemStatus *NodeStatusProperties `json:"system_status,omitempty"`
	MgmtClusterStatus *ManagementClusterNodeStatus `json:"mgmt_cluster_status,omitempty"`
	// Software version running on node
	Version string `json:"version,omitempty"`
	ControlClusterStatus *ControlClusterNodeStatus `json:"control_cluster_status,omitempty"`
}
