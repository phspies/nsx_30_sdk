package nsxt

// BGP routes per transport node.
type RoutesPerTransportNode struct {
	// Array of BGP neighbor route details for this transport node. 
	Routes []RouteDetails `json:"routes,omitempty"`
	// BGP neighbor source address.
	SourceAddress string `json:"source_address,omitempty"`
	// Transport node id
	TransportNodeId string `json:"transport_node_id,omitempty"`
}
