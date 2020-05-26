package nsxt

// A shaper that specifies ingress rate properties in kb/s
type IngressBroadcastRateShaper struct {
	Enabled bool `json:"enabled"`
	ResourceType string `json:"resource_type"`
	// Average bandwidth in kb/s
	AverageBandwidthKbps int32 `json:"average_bandwidth_kbps,omitempty"`
	// Peak bandwidth in kb/s
	PeakBandwidthKbps int32 `json:"peak_bandwidth_kbps,omitempty"`
	// Burst size in bytes
	BurstSizeBytes int32 `json:"burst_size_bytes,omitempty"`
}