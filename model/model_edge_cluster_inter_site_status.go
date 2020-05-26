package nsxt

type EdgeClusterInterSiteStatus struct {
	// Timestamp when the edge cluster inter-site status was last updated. 
	LastUpdateTimestamp int64 `json:"last_update_timestamp,omitempty"`
	// Overall status of all edge nodes IBGP status in the edge cluster. 
	OverallStatus string `json:"overall_status,omitempty"`
	// Name of the edge cluster whose status is being reported.
	EdgeClusterName string `json:"edge_cluster_name,omitempty"`
	// Per edge node inter-site status.
	MemberStatus []EdgeClusterMemberInterSiteStatus `json:"member_status,omitempty"`
	// Id of the edge cluster whose status is being reported.
	EdgeClusterId string `json:"edge_cluster_id,omitempty"`
}
