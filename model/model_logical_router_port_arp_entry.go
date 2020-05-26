package nsxt

type LogicalRouterPortArpEntry struct {
	// The IP address
	Ip string `json:"ip"`
	// The MAC address
	MacAddress string `json:"mac_address"`
}
