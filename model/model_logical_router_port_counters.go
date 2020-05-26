package nsxt

type LogicalRouterPortCounters struct {
	// The number of dropped packets
	DroppedPackets int64 `json:"dropped_packets,omitempty"`
	// The total number of bytes
	TotalBytes int64 `json:"total_bytes,omitempty"`
	// The total number of packets
	TotalPackets int64 `json:"total_packets,omitempty"`
}
