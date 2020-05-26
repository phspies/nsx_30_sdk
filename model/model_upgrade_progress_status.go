package nsxt

// Upgrade progress status
type UpgradeProgressStatus struct {
	// True if upgrade bundle is present
	UpgradeBundlePresent bool `json:"upgrade_bundle_present,omitempty"`
	// Status of last upgrade step
	LastUpgradeStepStatus *interface{} `json:"last_upgrade_step_status,omitempty"`
	// Meta info of upgrade
	UpgradeMetadata *interface{} `json:"upgrade_metadata,omitempty"`
}
