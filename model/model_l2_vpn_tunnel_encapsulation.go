package nsxt

// L2VPN tunnel encapsulation config
type L2VpnTunnelEncapsulation struct {
	// IP Address of the tunnel port. For hub, the IP is allocated from L2VpnService logical_tap_ip_pool. All sessions on same L2VpnService get the same local_endpoint_ip. For spoke, the IP must be provided.
	LocalEndpointIp string `json:"local_endpoint_ip,omitempty"`
	// Encapsulation protocol used by the tunnel
	Protocol string `json:"protocol,omitempty"`
	// IP Address of the peer tunnel port. For hub, the IP is allocated from L2VpnService logical_tap_ip_pool. For spoke, the IP must be provided.
	PeerEndpointIp string `json:"peer_endpoint_ip,omitempty"`
}
