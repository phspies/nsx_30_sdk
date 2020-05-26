package nsxt

// BGP neighbor learned/advertised route details.
type BgpNeighborRouteDetailsCsvRecord struct {
	// BGP Multi Exit Discriminator attribute.
	Med int64 `json:"med,omitempty"`
	// CIDR network address.
	Network string `json:"network,omitempty"`
	// BGP Weight attribute.
	Weight int64 `json:"weight,omitempty"`
	// Transport node id
	TransportNodeId string `json:"transport_node_id,omitempty"`
	// BGP AS path attribute.
	AsPath string `json:"as_path,omitempty"`
	// Next hop IP address.
	NextHop string `json:"next_hop,omitempty"`
	// Logical router id
	LogicalRouterId string `json:"logical_router_id,omitempty"`
	// BGP Local Preference attribute.
	LocalPref int64 `json:"local_pref,omitempty"`
	// BGP neighbor source address.
	SourceAddress string `json:"source_address,omitempty"`
	// BGP neighbor id
	NeighborId string `json:"neighbor_id,omitempty"`
	// BGP neighbor peer IP address.
	NeighborAddress string `json:"neighbor_address,omitempty"`
}
