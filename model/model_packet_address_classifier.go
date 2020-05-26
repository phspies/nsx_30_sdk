package nsxt

// A packet is classified to have an address binding, if its address configuration matches with all user specified properties. 
type PacketAddressClassifier struct {
	// Virtual Local Area Network Identifier
	Vlan int64 `json:"vlan,omitempty"`
	// A single IP address or a subnet, e.g. x.x.x.x or x.x.x.x/y
	IpAddress string `json:"ip_address,omitempty"`
	// A single MAC address
	MacAddress string `json:"mac_address,omitempty"`
}
