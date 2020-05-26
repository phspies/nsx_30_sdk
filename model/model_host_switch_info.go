package nsxt

// Information of host switch participating in transport zone
type HostSwitchInfo struct {
	// Type of a host switch
	HostSwitchType string `json:"host_switch_type,omitempty"`
	// Unique ID of a host switch
	HostSwitchId string `json:"host_switch_id,omitempty"`
	// Mode of host switch
	HostSwitchMode string `json:"host_switch_mode,omitempty"`
	// Name of a host switch
	HostSwitchName string `json:"host_switch_name,omitempty"`
}
