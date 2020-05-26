package nsxt

// Information of a network interface present on the partner appliance that needs to be configured by the NSX Manager.
type NicInfo struct {
	// Subnet mask associated with the NIC metadata.
	SubnetMask string `json:"subnet_mask,omitempty"`
	// Gateway address associated with the NIC metadata.
	GatewayAddress string `json:"gateway_address,omitempty"`
	// IP allocation type with values STATIC, DHCP, or NONE indicating that IP address is not required.
	IpAllocationType string `json:"ip_allocation_type,omitempty"`
	NicMetadata *NicMetadata `json:"nic_metadata,omitempty"`
	// Network Id associated with the NIC metadata. It can be a moref, or a logical switch ID. If it is to be taken from 'Agent VM Settings', then it should be empty.
	NetworkId string `json:"network_id,omitempty"`
	// If the nic should get IP using a static IP pool then IP pool id should be provided here.
	IpPoolId string `json:"ip_pool_id,omitempty"`
	// IP address associated with the NIC metadata. Required only when assigning IP statically for a deployment that is for a single VM instance.
	IpAddress string `json:"ip_address,omitempty"`
}
