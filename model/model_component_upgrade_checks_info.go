package nsxt

// Meta-data of pre/post-upgrade checks for a component
type ComponentUpgradeChecksInfo struct {
	// Collection of pre-upgrade checks
	PreUpgradeChecksInfo []UpgradeCheckInfo `json:"pre_upgrade_checks_info,omitempty"`
	// Collection of post-upgrade checks
	PostUpgradeChecksInfo []UpgradeCheckInfo `json:"post_upgrade_checks_info,omitempty"`
	// Component type of the pre/post-upgrade checks
	ComponentType string `json:"component_type"`
}
