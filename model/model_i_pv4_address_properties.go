package nsxt

// IPv4 address properties
type IPv4AddressProperties struct {
	// Interface netmask
	Netmask string `json:"netmask,omitempty"`
	// Interface IPv4 address
	IpAddress string `json:"ip_address,omitempty"`
}
