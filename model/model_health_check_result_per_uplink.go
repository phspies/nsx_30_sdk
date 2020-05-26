package nsxt

// Health check result for specific uplink. 
type HealthCheckResultPerUplink struct {
	// Name of the uplink.
	UplinkName string `json:"uplink_name,omitempty"`
	// List of VLAN ID ranges which are allowed by VLAN and MTU settings. 
	VlanAndMtuAllowed []HealthCheckVlanRange `json:"vlan_and_mtu_allowed,omitempty"`
	// List of VLAN ID ranges which may be disallowed by VLAN settings. 
	VlanDisallowed []HealthCheckVlanRange `json:"vlan_disallowed,omitempty"`
	// List of VLAN ID ranges which are allowed by VLAN settings but may be disallowed by MTU settings. 
	MtuDisallowed []HealthCheckVlanRange `json:"mtu_disallowed,omitempty"`
}
