package nsxt

// Ipaddress information of the fabric node.
type IpAddressInfo struct {
	// Source of the ipaddress information.
	Source string `json:"source,omitempty"`
	// IP Addresses of the the virtual network interface, as discovered in the source.
	IpAddresses []string `json:"ip_addresses,omitempty"`
}
