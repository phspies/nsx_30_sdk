package nsxt

// Provides IPSec VPN session status.
type IpSecVpnTransportStatus struct {
	// Resource types of L2VPN Transport tunnels
	ResourceType string `json:"resource_type"`
	TunnelId *ResourceReference `json:"tunnel_id,omitempty"`
	Status *IpSecVpnSessionStatus `json:"status,omitempty"`
}
