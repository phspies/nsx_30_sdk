package nsxt

type ContainerClusterStatus struct {
	// Display the container cluster status.
	Status string `json:"status,omitempty"`
	// Display the cluster check interval in seconds.
	Interval int32 `json:"interval,omitempty"`
	// Identifier of the container cluster.
	ClusterId string `json:"cluster_id,omitempty"`
	// Detail information on status.
	Detail string `json:"detail,omitempty"`
}
