package nsxt

// DHCP option 121 to define classless static route.
type DhcpOption121 struct {
	// Classless static route of DHCP option 121.
	StaticRoutes []ClasslessStaticRoute `json:"static_routes"`
}
