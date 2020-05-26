package nsxt

type DhcpLeasePerIp struct {
	// subnet of client network
	Subnet string `json:"subnet,omitempty"`
	// lease time of the ip address, in seconds
	LeaseTime string `json:"lease_time,omitempty"`
	// mac address of client
	MacAddress string `json:"mac_address"`
	// expire time of the lease
	ExpireTime string `json:"expire_time,omitempty"`
	// start time of lease
	StartTime string `json:"start_time"`
	// ip address of client
	IpAddress string `json:"ip_address"`
}
