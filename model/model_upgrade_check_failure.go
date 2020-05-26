package nsxt

// Pre/post-upgrade check failure
type UpgradeCheckFailure struct {
	// Component type of the origin of failure
	ComponentType string `json:"component_type,omitempty"`
	// Unique id of origin of pre/post-upgrade check failure
	OriginId string `json:"origin_id,omitempty"`
	Message *UpgradeCheckFailureMessage `json:"message,omitempty"`
	// Name of origin of pre/post-upgrade check failure
	OriginName string `json:"origin_name,omitempty"`
	// Type of the pre/post-upgrade check failure
	Type_ string `json:"type,omitempty"`
	// Type of origin of pre/post-upgrade check failure
	OriginType string `json:"origin_type,omitempty"`
}
