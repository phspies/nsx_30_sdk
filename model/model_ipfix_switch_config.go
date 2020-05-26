package nsxt

// This is deprecated. Please use IpfixSwitchUpmProfile instead which can specify its own collectors and observation ID. 
type IpfixSwitchConfig struct {
	// List of objects where the IPFIX Config will be enabled.
	AppliedTos []ResourceReference `json:"applied_tos,omitempty"`
	// Supported IPFIX Config Types.
	ResourceType string `json:"resource_type"`
	// The time in seconds after a Flow is expired even if more packets matching this Flow are received by the cache. 
	ActiveTimeout int32 `json:"active_timeout,omitempty"`
	// The time in seconds after a Flow is expired if no more packets matching this Flow are received by the cache. 
	IdleTimeout int32 `json:"idle_timeout,omitempty"`
	// The probability in percentage that a packet is sampled. The value should be  in range (0,100] and can only have three decimal places at most. The probability  is equal for every packet. 
	PacketSampleProbability float64 `json:"packet_sample_probability,omitempty"`
	// The maximum number of flow entries in each exporter flow cache. 
	MaxFlows int64 `json:"max_flows,omitempty"`
}
