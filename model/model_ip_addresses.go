package nsxt

type IpAddresses struct {
	// The IP addresses in the form of IP Address, IP Range, CIDR, used as source IPs or destination IPs of filters.
	IpAddresses []string `json:"ip_addresses,omitempty"`
}
