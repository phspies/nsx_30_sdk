package nsxt

// IKE session traffic summary provides IKE session status and aggregate of traffic across all tunnel.
type IpSecVpnSessionStatus struct {
	// UUID of vpn session.
	IpsecVpnSessionId string `json:"ipsec_vpn_session_id,omitempty"`
	// Display name of vpn session.
	DisplayName string `json:"display_name,omitempty"`
	// Number of failed tunnels.
	FailedTunnels int64 `json:"failed_tunnels,omitempty"`
	// Number of negotiated tunnels.
	NegotiatedTunnels int64 `json:"negotiated_tunnels,omitempty"`
	// Gives session status consolidated using IKE status and tunnel status. It can be UP, DOWN, DEGRADED. If IKE and all tunnels are UP status will be UP, if all down it will be DOWN, otherwise it will be DEGRADED.
	SessionStatus string `json:"session_status,omitempty"`
	// Timestamp when the data was last updated.
	LastUpdateTimestamp int64 `json:"last_update_timestamp,omitempty"`
	AggregateTrafficCounters *IpSecVpnTrafficCounters `json:"aggregate_traffic_counters,omitempty"`
	IkeStatus *IpSecVpnikeSessionStatus `json:"ike_status,omitempty"`
	// Total number of tunnels.
	TotalTunnels int64 `json:"total_tunnels,omitempty"`
}
