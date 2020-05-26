package nsxt

type IpInfo struct {
	// IPv4 Addresses
	IpAddresses []string `json:"ip_addresses"`
	// Subnet Prefix Length
	PrefixLength int64 `json:"prefix_length"`
}
