package nsxt

type EdgeClusterMemberStatus struct {
	TransportNode *ResourceReference `json:"transport_node"`
	// Status of an edge node
	Status string `json:"status"`
}
