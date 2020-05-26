package nsxt

// VlanTrunkspec is used for specifying trunk VLAN id ranges.
type VlanTrunkSpec struct {
	// Trunk VLAN id ranges
	VlanRanges []TrunkVlanRange `json:"vlan_ranges"`
}
