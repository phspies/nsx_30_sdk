package nsxt

// The configuration of the NSX cluster. The cluster configuration consists of a list of cluster node attributes.
type ClusterConfiguration struct {
	// Cluster configuration version
	ConfigVersion int64 `json:"config_version,omitempty"`
	// Nodes in the cluster configuration
	Nodes []ClusterNode `json:"nodes,omitempty"`
	// UUID of the cluster
	ClusterId string `json:"cluster_id,omitempty"`
}
