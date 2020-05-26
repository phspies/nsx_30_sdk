package nsxt

// Transport tunnel status.
type L2VpnTransportTunnelStatus struct {
	// Resource types of L2VPN Transport tunnels
	ResourceType string `json:"resource_type"`
	TunnelId *ResourceReference `json:"tunnel_id,omitempty"`
}
