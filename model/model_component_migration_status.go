package nsxt

type ComponentMigrationStatus struct {
	// Migration status of component
	Status string `json:"status,omitempty"`
	// Indicator of migration progress in percentage
	PercentComplete float64 `json:"percent_complete,omitempty"`
	// Can the migration of the remaining units in this component be skipped
	CanSkip bool `json:"can_skip,omitempty"`
	// Details about the migration status
	Details string `json:"details,omitempty"`
	// Component type for the migration status
	ComponentType string `json:"component_type,omitempty"`
}
