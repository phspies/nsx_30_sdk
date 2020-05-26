package nsxt

type UpgradeUnitTypeStats struct {
	// Number of nodes with issues that may cause upgrade failure
	NodeWithIssuesCount int32 `json:"node_with_issues_count,omitempty"`
	// Number of nodes
	NodeCount int32 `json:"node_count,omitempty"`
	// Version of the upgrade unit
	Version string `json:"version,omitempty"`
	// Type of upgrade unit
	Type_ string `json:"type,omitempty"`
}
