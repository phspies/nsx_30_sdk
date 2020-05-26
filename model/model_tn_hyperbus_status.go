package nsxt

type TnHyperbusStatus struct {
	// Display the hyperbus status
	HyperbusStatus string `json:"hyperbus_status"`
	// Transport node id.
	TransportNodeId string `json:"transport_node_id"`
}
