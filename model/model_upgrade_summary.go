package nsxt

type UpgradeSummary struct {
	// Has upgrade coordinator been updated after upload of upgrade bundle file
	UpgradeCoordinatorUpdated bool `json:"upgrade_coordinator_updated,omitempty"`
	// Target system version
	TargetVersion string `json:"target_version,omitempty"`
	// Current version of upgrade coordinator
	UpgradeCoordinatorVersion string `json:"upgrade_coordinator_version,omitempty"`
	// Status of upgrade
	UpgradeStatus string `json:"upgrade_status,omitempty"`
	ComponentTargetVersions []ComponentTargetVersion `json:"component_target_versions,omitempty"`
	// Current system version
	SystemVersion string `json:"system_version,omitempty"`
	// Name of the last successfully uploaded upgrade bundle file
	UpgradeBundleFileName string `json:"upgrade_bundle_file_name,omitempty"`
}
