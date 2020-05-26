package nsxt

// IPSec VPN policy traffic statistics
type IpSecVpnPolicyTrafficStatistics struct {
	AggregateTrafficCounters *IpSecVpnTrafficCounters `json:"aggregate_traffic_counters,omitempty"`
	// Tunnel statistics.
	TunnelStatistics []IpSecVpnTunnelTrafficStatistics `json:"tunnel_statistics,omitempty"`
	// Tunnel port identifier.
	TunnelPortId string `json:"tunnel_port_id,omitempty"`
	// Policy Identifier.
	PolicyId string `json:"policy_id,omitempty"`
}
