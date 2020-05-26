package nsxt

// DHCP static binding to define a static ip allocation.
type DhcpStaticBinding struct {
	// Lease time, in seconds, [60-(2^32-1)]. Default is 86400.
	LeaseTime int64 `json:"lease_time,omitempty"`
	// Gateway ip address of the allocation.
	GatewayIp string `json:"gateway_ip,omitempty"`
	Options *DhcpOptions `json:"options,omitempty"`
	// The ip address to be assigned to the host.
	IpAddress string `json:"ip_address"`
	// The host name to be assigned to the host.
	HostName string `json:"host_name,omitempty"`
	// The MAC address of the host.
	MacAddress string `json:"mac_address"`
}
