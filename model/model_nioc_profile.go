package nsxt

// This profile is created for Network I/O Control(NIOC). 
type NiocProfile struct {
	RequiredCapabilities []string `json:"required_capabilities,omitempty"`
	// Supported HostSwitch profiles.
	ResourceType string `json:"resource_type"`
	// host_infra_traffic_res specifies bandwidth allocation for various traffic resources. 
	HostInfraTrafficRes []ResourceAllocation `json:"host_infra_traffic_res,omitempty"`
	// The enabled property specifies the status of NIOC feature. When enabled is set to true, NIOC feature is turned on and the bandwidth allocations specified for the traffic resources are enforced. When enabled is set to false, NIOC feature is turned off and no bandwidth allocation is guaranteed. By default, enabled will be set to true. 
	Enabled bool `json:"enabled,omitempty"`
}
