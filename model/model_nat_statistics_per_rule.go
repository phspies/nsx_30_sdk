package nsxt

type NatStatisticsPerRule struct {
	// The number of packets
	TotalPackets int64 `json:"total_packets,omitempty"`
	// The number of bytes
	TotalBytes int64 `json:"total_bytes,omitempty"`
	// The number of active sessions
	ActiveSessions int64 `json:"active_sessions,omitempty"`
	// Timestamp when the data was last updated; unset if data source has never updated the data.
	LastUpdateTimestamp int64 `json:"last_update_timestamp,omitempty"`
	// The id of the logical router which owns the NAT rule.
	LogicalRouterId string `json:"logical_router_id,omitempty"`
	// The id of the NAT rule.
	Id string `json:"id,omitempty"`
	// The warning message about the NAT Rule statistics.
	WarningMessage string `json:"warning_message,omitempty"`
}
