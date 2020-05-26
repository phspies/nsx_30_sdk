package nsxt

// IP and MAC pair.
type IpMacPair struct {
	// IP address
	Ip string `json:"ip"`
	// MAC address
	Mac string `json:"mac,omitempty"`
}
