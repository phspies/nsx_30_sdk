package nsxt

type EdgeClusterStatus struct {
	// Timestamp when the cluster status was last updated
	LastUpdateTimestamp int64 `json:"last_update_timestamp,omitempty"`
	// Status of an edge node
	EdgeClusterStatus string `json:"edge_cluster_status"`
	// Per Edge Node Status
	MemberStatus []EdgeClusterMemberStatus `json:"member_status,omitempty"`
	// Id of the edge cluster whose status is being reported
	EdgeClusterId string `json:"edge_cluster_id"`
}
