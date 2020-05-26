package nsxt

// A Route Based VPN is more flexible, more powerful and recommended over policy based VPN. IP Tunnel port is created and all traffic routed via tunnel port is protected. Routes can be configured statically or can be learned through BGP. A route based VPN is must for establishing redundant VPN session to remote site.
type RouteBasedIpSecVpnSession struct {
	// Peer endpoint identifier.
	PeerEndpointId string `json:"peer_endpoint_id"`
	// Identifier of VPN Service linked with local endpoint.
	IpsecVpnServiceId string `json:"ipsec_vpn_service_id,omitempty"`
	// Local endpoint identifier.
	LocalEndpointId string `json:"local_endpoint_id"`
	TcpMssClamping *TcpMssClamping `json:"tcp_mss_clamping,omitempty"`
	// Enable/Disable IPSec VPN session.
	Enabled bool `json:"enabled,omitempty"`
	// A Policy Based VPN requires to define protect rules that match   local and peer subnets. IPSec security associations is   negotiated for each pair of local and peer subnet. A Route Based VPN is more flexible, more powerful and recommended over   policy based VPN. IP Tunnel port is created and all traffic routed via   tunnel port is protected. Routes can be configured statically   or can be learned through BGP. A route based VPN is must for establishing   redundant VPN session to remote site. 
	ResourceType string `json:"resource_type"`
	// IP Tunnel ports.
	TunnelPorts []TunnelPortConfig `json:"tunnel_ports"`
}
