package nsxt

// Profile for BFD health monitoring
type BfdHealthMonitoringProfile struct {
	ResourceType string `json:"resource_type"`
	// The time interval (in millisec) between probe packets for tunnels between transport nodes.
	ProbeInterval int64 `json:"probe_interval,omitempty"`
	// The flag is to turn on/off latency. A POST or PUT request with \"latency_enabled\" true will enable NSX to send the networking latency data to thrid-party monitoring tools like vRNI.
	LatencyEnabled bool `json:"latency_enabled,omitempty"`
	// Whether the heartbeat is enabled. A POST or PUT request with \"enabled\" false (with no probe intervals) will set (POST) or reset (PUT) the probe_interval to their default value.
	Enabled bool `json:"enabled"`
}
