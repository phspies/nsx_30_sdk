package nsxt

// Trunk VLAN id range
type TrunkVlanRange struct {
	// Virtual Local Area Network Identifier
	Start int64 `json:"start"`
	// Virtual Local Area Network Identifier
	End int64 `json:"end"`
}
