package nsxt

// IP Tunnel port configuration.
type TunnelPortConfig struct {
	// IP Tunnel port  (commonly referred as VTI) subnet.
	IpSubnets []IpSubnet `json:"ip_subnets"`
	// Logical route port identifier.
	TunnelPortId string `json:"tunnel_port_id,omitempty"`
}
