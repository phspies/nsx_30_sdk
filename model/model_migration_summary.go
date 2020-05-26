package nsxt

type MigrationSummary struct {
	// Target system version
	TargetVersion string `json:"target_version,omitempty"`
	// Status of migration
	MigrationStatus string `json:"migration_status,omitempty"`
	// Current version of migration coordinator
	MigrationCoordinatorVersion string `json:"migration_coordinator_version,omitempty"`
	// Current system version
	SystemVersion string `json:"system_version,omitempty"`
	ComponentTargetVersions []ComponentTargetVersion `json:"component_target_versions,omitempty"`
}
