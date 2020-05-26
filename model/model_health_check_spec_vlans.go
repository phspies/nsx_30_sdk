package nsxt

// HealthCheckSpecVlan is used for specifying VLAN ID ranges for healthcheck. 
type HealthCheckSpecVlans struct {
	// VLAN ID ranges
	VlanRanges []HealthCheckVlanRange `json:"vlan_ranges"`
}
