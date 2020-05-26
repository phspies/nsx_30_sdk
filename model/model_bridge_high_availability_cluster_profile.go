package nsxt

// Profile for BFD HA cluster setting
type BridgeHighAvailabilityClusterProfile struct {
	// Supported cluster profiles.
	ResourceType string `json:"resource_type"`
	// whether the heartbeat is enabled
	Enable bool `json:"enable,omitempty"`
	// the time interval (in millisec) between probe packets for heartbeat purpose
	BfdProbeInterval int64 `json:"bfd_probe_interval,omitempty"`
}
