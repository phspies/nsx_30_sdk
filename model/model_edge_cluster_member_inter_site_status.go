package nsxt

type EdgeClusterMemberInterSiteStatus struct {
	TransportNode *ResourceReference `json:"transport_node,omitempty"`
	// Total number of current established inter-site IBGP sessions.
	EstablishedBgpSessions int64 `json:"established_bgp_sessions,omitempty"`
	// Edge node IBGP status
	Status string `json:"status,omitempty"`
	// Total number of inter-site IBGP sessions.
	TotalBgpSessions int64 `json:"total_bgp_sessions,omitempty"`
	// Inter-site BGP neighbor status.
	NeighborStatus []BgpNeighborStatusLiteDto `json:"neighbor_status,omitempty"`
}
