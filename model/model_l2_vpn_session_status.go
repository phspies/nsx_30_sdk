package nsxt

// L2VPN session status.
type L2VpnSessionStatus struct {
	// L2 VPN session status, specifies UP/DOWN.
	Status string `json:"status,omitempty"`
	// Transport tunnels status.
	TransportTunnels []L2VpnTransportTunnelStatus `json:"transport_tunnels,omitempty"`
	// L2VPN display name.
	DisplayName string `json:"display_name,omitempty"`
	// L2VPN session identifier.
	SessionId string `json:"session_id,omitempty"`
}
