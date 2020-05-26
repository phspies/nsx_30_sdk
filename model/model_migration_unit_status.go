package nsxt

type MigrationUnitStatus struct {
	// Status of migration unit
	Status string `json:"status,omitempty"`
	// List of errors occurred during migration of this migration unit
	Errors []string `json:"errors,omitempty"`
	// Name of migration unit
	DisplayName string `json:"display_name,omitempty"`
	// Identifier of migration unit
	Id string `json:"id,omitempty"`
	// Indicator of migration progress in percentage
	PercentComplete float64 `json:"percent_complete,omitempty"`
}
