package nsxt

type ClustersAggregateInfo struct {
	// Array of Management Nodes
	ManagementCluster []ManagementNodeAggregateInfo `json:"management_cluster"`
	// Array of Controller Nodes
	ControllerCluster []ControllerNodeAggregateInfo `json:"controller_cluster"`
	ClusterStatus *AllClusterGroupStatus `json:"cluster_status,omitempty"`
}
