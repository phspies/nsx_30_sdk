package nsxt

// VLAN ID range
type HealthCheckVlanRange struct {
	// Virtual Local Area Network Identifier
	Start int64 `json:"start"`
	// Virtual Local Area Network Identifier
	End int64 `json:"end"`
}
