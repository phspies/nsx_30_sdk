package nsxt

type MigrationUnitGroupStatus struct {
	// Migration status of migration unit group
	Status string `json:"status,omitempty"`
	// Number of nodes in the migration unit group that failed migration
	FailedCount int32 `json:"failed_count,omitempty"`
	// Number of migration units in the group
	MigrationUnitCount int32 `json:"migration_unit_count,omitempty"`
	// Identifier for migration unit group
	GroupId string `json:"group_id,omitempty"`
	// Indicator of migration progress in percentage
	PercentComplete float64 `json:"percent_complete,omitempty"`
	// Name of the migration unit group
	GroupName string `json:"group_name,omitempty"`
}
