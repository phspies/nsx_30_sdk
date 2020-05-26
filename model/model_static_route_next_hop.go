package nsxt

type StaticRouteNextHop struct {
	// Action to be taken on matching packets for NULL routes.
	BlackholeAction string `json:"blackhole_action,omitempty"`
	// Administrative Distance for the next hop IP
	AdministrativeDistance int64 `json:"administrative_distance,omitempty"`
	// Next Hop IP
	IpAddress string `json:"ip_address,omitempty"`
	// Status of bfd for this next hop where bfd_enabled = true indicate bfd is enabled for this next hop and bfd_enabled = false indicate bfd peer is disabled or not configured for this next hop.
	BfdEnabled bool `json:"bfd_enabled,omitempty"`
	LogicalRouterPortId *ResourceReference `json:"logical_router_port_id,omitempty"`
}
