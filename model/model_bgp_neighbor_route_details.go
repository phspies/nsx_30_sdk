package nsxt

// BGP neighbor learned/advertised route details.
type BgpNeighborRouteDetails struct {
	// Logical router id
	LogicalRouterId string `json:"logical_router_id,omitempty"`
	// BGP neighbor id
	NeighborId string `json:"neighbor_id,omitempty"`
	// Array of BGP neighbor route details per transport node. 
	PerTransportNodeRoutes []RoutesPerTransportNode `json:"per_transport_node_routes,omitempty"`
	// BGP neighbor peer IP address.
	NeighborAddress string `json:"neighbor_address,omitempty"`
}
