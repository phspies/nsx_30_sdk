package nsxt

type MigrationUnitAggregateInfo struct {
	Self *SelfResourceLink `json:"_self,omitempty"`
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink `json:"_links,omitempty"`
	// Schema for this resource
	Schema string `json:"_schema,omitempty"`
	// Status of migration unit
	Status string `json:"status,omitempty"`
	// Indicator of migration progress in percentage
	PercentComplete float64 `json:"percent_complete,omitempty"`
	// List of errors occurred during migration of this migration unit
	Errors []string `json:"errors,omitempty"`
	Unit *MigrationUnit `json:"unit,omitempty"`
}
