package nsxt

// The type provides the information of a non-running cluster node required for the initialization of a management cluster. The administrator needs to start this node for management cluster to initialize properly (or decommission it explicitly).
type ClusterInitializationNodeInfo struct {
	// The (internal) disk-store ID of the member
	DiskStoreId string `json:"disk_store_id,omitempty"`
	// The IP address (or domain name) of the cluster node
	HostAddress string `json:"host_address,omitempty"`
}
