package nsxt

type MigrationUnitTypeStats struct {
	// Number of nodes with issues that may cause migration failure
	NodeWithIssuesCount int32 `json:"node_with_issues_count,omitempty"`
	// Number of nodes
	NodeCount int32 `json:"node_count,omitempty"`
	// Version of the migration unit
	Version string `json:"version,omitempty"`
	// Type of migration unit
	Type_ string `json:"type,omitempty"`
}
