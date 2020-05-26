package nsxt

// Session statistics gives aggregated statistics of all policies for all the tunnels.
type IpSecVpnSessionStatistics struct {
	IkeTrafficStatistics *IpSecVpnikeTrafficStatistics `json:"ike_traffic_statistics,omitempty"`
	// Display name of vpn session.
	DisplayName string `json:"display_name,omitempty"`
	// Gives aggregate traffic statistics across all ipsec tunnels and individual tunnel statistics.
	PolicyStatistics []IpSecVpnPolicyTrafficStatistics `json:"policy_statistics,omitempty"`
	// Partial statistics if true specifies that the statistics are only from active node.
	PartialStats bool `json:"partial_stats,omitempty"`
	// UUID of vpn session.
	IpsecVpnSessionId string `json:"ipsec_vpn_session_id,omitempty"`
	// Timestamp when the data was last updated.
	LastUpdateTimestamp int64 `json:"last_update_timestamp,omitempty"`
	IkeStatus *IpSecVpnikeSessionStatus `json:"ike_status,omitempty"`
	AggregateTrafficCounters *IpSecVpnTrafficCounters `json:"aggregate_traffic_counters,omitempty"`
}
