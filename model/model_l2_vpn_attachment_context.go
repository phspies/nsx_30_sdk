package nsxt

type L2VpnAttachmentContext struct {
	// A flag to indicate whether to allocate addresses from allocation     pools bound to the parent logical switch. 
	AllocateAddresses string `json:"allocate_addresses,omitempty"`
	// Used to identify which concrete class it is
	ResourceType string `json:"resource_type"`
	// List of local egress IP addresses, used for local egress optimization. 
	LocalEgressIp []string `json:"local_egress_ip,omitempty"`
	// Tunnel Id to uniquely identify the extension.
	TunnelId int32 `json:"tunnel_id"`
}
