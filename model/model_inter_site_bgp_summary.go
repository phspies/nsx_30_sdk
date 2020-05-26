package nsxt

type InterSiteBgpSummary struct {
	// Timestamp when the inter-site IBgp neighbors status was last updated. 
	LastUpdateTimestamp int64 `json:"last_update_timestamp,omitempty"`
	// Status of all inter-site IBgp neighbors.
	NeighborStatus []BgpNeighborStatus `json:"neighbor_status,omitempty"`
	// Edge node id whose status is being reported.
	TransportNodeId string `json:"transport_node_id,omitempty"`
}
