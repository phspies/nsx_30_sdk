package nsxt

// NSX global configs for operation collector
type OperationCollectorGlobalConfig struct {
	// Valid Global configuration types
	ResourceType string `json:"resource_type"`
	// Report interval for operation data in seconds.
	ReportInterval int32 `json:"report_interval,omitempty"`
	// Operation Collector Config.
	Collectors []OperationCollector `json:"collectors,omitempty"`
}
