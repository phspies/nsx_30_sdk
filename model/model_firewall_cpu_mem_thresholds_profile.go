package nsxt

// A profile holding CPU and memory threshold configuration.
type FirewallCpuMemThresholdsProfile struct {
	// Resource type to use as profile type
	ResourceType string `json:"resource_type"`
	// Heap memory threshold percentage to monitor and report for distributed firewall.
	MemThresholdPercentage int64 `json:"mem_threshold_percentage"`
	// CPU utilization threshold percentage to monitor and report for distributed firewall.
	CpuThresholdPercentage int64 `json:"cpu_threshold_percentage"`
}
