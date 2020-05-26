package nsxt

// DHCP classless static route option.
type ClasslessStaticRoute struct {
	// IP address of next hop of the route.
	NextHop string `json:"next_hop"`
	// Destination network in CIDR format.
	Network string `json:"network"`
}
