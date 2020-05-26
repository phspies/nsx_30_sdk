package nsxt

// TAP (Terminal access point) traffic statistics for L2VPN.
type L2VpnTapTrafficStatistics struct {
	// Total number of outgoing packets.
	PacketsOut int64 `json:"packets_out,omitempty"`
	// Total number of incoming bytes.
	BytesIn int64 `json:"bytes_in,omitempty"`
	// Total number of packets dropped while sending for any reason.
	PacketsSentError int64 `json:"packets_sent_error,omitempty"`
	// Total number of incoming packets dropped.
	PacketsReceiveError int64 `json:"packets_receive_error,omitempty"`
	// Total number of incoming packets.
	PacketsIn int64 `json:"packets_in,omitempty"`
	// Total number of outgoing bytes.
	BytesOut int64 `json:"bytes_out,omitempty"`
}
