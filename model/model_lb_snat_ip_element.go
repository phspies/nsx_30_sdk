package nsxt

type LbSnatIpElement struct {
	// Ip address or ip range such as 1.1.1.1 or 1.1.1.101-1.1.1.160
	IpAddress string `json:"ip_address"`
	// Subnet prefix length should be not specified if there is only one single IP address or IP range. 
	PrefixLength int64 `json:"prefix_length,omitempty"`
}
