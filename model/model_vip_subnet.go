package nsxt

type VipSubnet struct {
	// Subnet Prefix Length
	PrefixLength int64 `json:"prefix_length"`
	// Array of IP address subnets which will be used as floating IP addresses. | These IPs will move and will be owned by Active node.
	ActiveVipAddresses []string `json:"active_vip_addresses"`
}
