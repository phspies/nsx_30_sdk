package nsxt

type NatCounters struct {
	// The number of packets
	TotalPackets int64 `json:"total_packets,omitempty"`
	// The number of bytes
	TotalBytes int64 `json:"total_bytes,omitempty"`
	// The number of active sessions
	ActiveSessions int64 `json:"active_sessions,omitempty"`
}
