package nsxt

// Meta-data of a pre/post-upgrade check
type UpgradeCheckInfo struct {
	// Display name of the pre/post-upgrade check
	Name string `json:"name,omitempty"`
	// Component type of the pre/post-upgrade check
	ComponentType string `json:"component_type"`
	// Description of the pre/post-upgrade check
	Description string `json:"description,omitempty"`
}
