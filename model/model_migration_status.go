package nsxt

type MigrationStatus struct {
	// Status of migration
	OverallMigrationStatus string `json:"overall_migration_status,omitempty"`
	// List of component statuses
	ComponentStatus []ComponentMigrationStatus `json:"component_status,omitempty"`
}
