package nsxt

// An IP prefix to mark the private network that NSX-Intelligence network flow is collected from. 
type IntelligenceFlowPrivateIpPrefixInfo struct {
	// The type of IP address. 
	AddressType string `json:"address_type"`
	// The length of IP address prefix that marks the range of private network. 
	PrefixLength int64 `json:"prefix_length"`
	// The prefix of IP address that marks the range of private network. 
	Address string `json:"address"`
}
