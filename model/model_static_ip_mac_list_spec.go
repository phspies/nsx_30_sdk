package nsxt

// IP and MAC assignment specification for Static IP List.
type StaticIpMacListSpec struct {
	ResourceType string `json:"resource_type"`
	// Subnet mask
	SubnetMask string `json:"subnet_mask"`
	// List of IPs and MACs for transport node host switch virtual tunnel endpoints
	IpMacList []IpMacPair `json:"ip_mac_list"`
	// Gateway IP
	DefaultGateway string `json:"default_gateway"`
}
