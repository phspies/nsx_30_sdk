package nsxt

// A shaper that specifies egress rate properties in Mb/s
type EgressRateShaper struct {
	Enabled bool `json:"enabled"`
	ResourceType string `json:"resource_type"`
	// Average bandwidth in Mb/s
	AverageBandwidthMbps int32 `json:"average_bandwidth_mbps,omitempty"`
	// Peak bandwidth in Mb/s
	PeakBandwidthMbps int32 `json:"peak_bandwidth_mbps,omitempty"`
	// Burst size in bytes
	BurstSizeBytes int32 `json:"burst_size_bytes,omitempty"`
}
