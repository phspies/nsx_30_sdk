package nsxt

type VlanMirrorSource struct {
	// Resource types of mirror source
	ResourceType string `json:"resource_type"`
	// Source VLAN ID list
	VlanIds []int64 `json:"vlan_ids"`
}
