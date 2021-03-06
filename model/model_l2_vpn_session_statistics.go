package nsxt

// Session statistics gives VPN session status and traffic statistics per logical switch.
type L2VpnSessionStatistics struct {
	// Traffic statistics per logical switch.
	TrafficStatisticsPerLogicalSwitch []L2VpnPerLsTrafficStatistics `json:"traffic_statistics_per_logical_switch,omitempty"`
	// L2VPN display name.
	DisplayName string `json:"display_name,omitempty"`
	// Partial statistics is set to true if onle active node responds while standby does not. In case of both nodes responded statistics will be summed and partial stats will be false. If cluster has only active node, partial statistics will always be false.
	PartialStats bool `json:"partial_stats,omitempty"`
	// Session identifier for L2VPN.
	SessionId string `json:"session_id,omitempty"`
	// Tunnel port traffic counters.
	TapTrafficCounters []L2VpnTapTrafficStatistics `json:"tap_traffic_counters,omitempty"`
}
