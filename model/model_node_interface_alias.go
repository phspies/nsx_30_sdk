package nsxt

// Node network interface alias
type NodeInterfaceAlias struct {
	// Interface configuration
	IpConfiguration string `json:"ip_configuration,omitempty"`
	// Interface netmask
	Netmask string `json:"netmask,omitempty"`
	// Interface IP address
	IpAddress string `json:"ip_address,omitempty"`
	// Interface MAC address
	PhysicalAddress string `json:"physical_address,omitempty"`
	// Interface broadcast address
	BroadcastAddress string `json:"broadcast_address,omitempty"`
}
