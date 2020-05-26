package nsxt

// A list of the statuses of all the groups in the cluster.
type AllClusterGroupStatus struct {
	// Overall status of the cluster
	OverallStatus string `json:"overall_status,omitempty"`
	// UUID of the cluster
	ClusterId string `json:"cluster_id,omitempty"`
	// Array of groups and their statuses
	Groups []ClusterGroupStatus `json:"groups,omitempty"`
}
