package nsxt

type LogicalRouterRouteEntry struct {
	// Logical router component(Service Router/Distributed Router) id
	LrComponentId string `json:"lr_component_id,omitempty"`
	// The IP address of the next hop
	NextHop string `json:"next_hop,omitempty"`
	// Logical router component(Service Router/Distributed Router) type
	LrComponentType string `json:"lr_component_type,omitempty"`
	// CIDR network address
	Network string `json:"network"`
	// Route type (USER, CONNECTED, NSX_INTERNAL,..)
	RouteType string `json:"route_type"`
	// The id of the logical router port which is used as the next hop
	LogicalRouterPortId string `json:"logical_router_port_id,omitempty"`
	// The admin distance of the next hop
	AdminDistance int64 `json:"admin_distance,omitempty"`
}
