package nsxt

// Health check result on specific host switch of specific transport node. 
type HealthCheckResultOnHostSwitch struct {
	// Status of VLAN-MTU health check result on host switch. 
	VlanMtuStatus string `json:"vlan_mtu_status,omitempty"`
	// Name of the host switch.
	HostSwitchName string `json:"host_switch_name,omitempty"`
	// List of health check results per uplink on current host switch of specific transport node. 
	ResultsPerUplink []HealthCheckResultPerUplink `json:"results_per_uplink,omitempty"`
	// Timestamp of check result updated.
	UpdatedTime int64 `json:"updated_time,omitempty"`
}
