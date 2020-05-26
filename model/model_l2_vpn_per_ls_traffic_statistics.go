package nsxt

// Traffic statistics for a logical switch.
type L2VpnPerLsTrafficStatistics struct {
	// Total number of outgoing packets.
	PacketsOut int64 `json:"packets_out,omitempty"`
	// Total number of incoming packets dropped.
	PacketsReceiveError int64 `json:"packets_receive_error,omitempty"`
	// Total number of incoming bytes.
	BytesIn int64 `json:"bytes_in,omitempty"`
	// Total number of incoming Broadcast, Unknown unicast and Multicast (BUM) packets.
	BumPacketsIn int64 `json:"bum_packets_in,omitempty"`
	// Total number of outgoing Broadcast, Unknown unicast and Multicast (BUM) bytes.
	BumBytesOut int64 `json:"bum_bytes_out,omitempty"`
	LogicalSwitch *ResourceReference `json:"logical_switch,omitempty"`
	// Total number of outgoing bytes.
	BytesOut int64 `json:"bytes_out,omitempty"`
	// Total number of packets dropped while sending for any reason.
	PacketsSentError int64 `json:"packets_sent_error,omitempty"`
	// Total number of outgoing Broadcast, Unknown unicast and Multicast (BUM) packets.
	BumPacketsOut int64 `json:"bum_packets_out,omitempty"`
	// Total number of incoming packets.
	PacketsIn int64 `json:"packets_in,omitempty"`
	// Total number of incoming Broadcast, Unknown unicast and Multicast (BUM) bytes.
	BumBytesIn int64 `json:"bum_bytes_in,omitempty"`
}
