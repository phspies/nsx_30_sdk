package nsxt

// Distributed virtual portgroup on a VC
type DistributedVirtualPortgroup struct {
	// External id of the virtual portgroup
	ExternalId string `json:"external_id,omitempty"`
	// Portgroup type like DistributedVirtualPortgroup
	OriginType string `json:"origin_type,omitempty"`
	// Id of the portgroup, eg. a mo-ref from VC.
	CmLocalId string `json:"cm_local_id,omitempty"`
	// This parameters reflects the managed entity status of the portgroup as reported by VC. 
	OverallStatus string `json:"overall_status,omitempty"`
	// For distributed virtual portgroup, backing type is standard. For logical switch portgroup, the backing type is set to nsx. 
	BackingType string `json:"backing_type,omitempty"`
	// Generated UUID of the portgroup
	Key string `json:"key,omitempty"`
}
