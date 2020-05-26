package nsxt

type UpgradeUnitAggregateInfo struct {
	Self *SelfResourceLink `json:"_self,omitempty"`
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink `json:"_links,omitempty"`
	// Schema for this resource
	Schema string `json:"_schema,omitempty"`
	// Status of upgrade unit
	Status string `json:"status,omitempty"`
	PreUpgradeChecks *UpgradeCheckListResults `json:"pre_upgrade_checks,omitempty"`
	// List of errors occurred during upgrade of this upgrade unit
	Errors []string `json:"errors,omitempty"`
	// Name of the upgrade unit
	DisplayName string `json:"display_name,omitempty"`
	PostUpgradeChecks *UpgradeCheckListResults `json:"post_upgrade_checks,omitempty"`
	// List of warnings indicating issues with the upgrade unit that may result in upgrade failure
	Warnings []string `json:"warnings,omitempty"`
	// This is component version e.g. if upgrade unit is of type edge, then this is edge version.
	CurrentVersion string `json:"current_version,omitempty"`
	Group *UpgradeUnitGroupInfo `json:"group,omitempty"`
	// Indicator of upgrade progress in percentage
	PercentComplete float64 `json:"percent_complete,omitempty"`
	// Upgrade unit type
	Type_ string `json:"type,omitempty"`
	// Identifier of the upgrade unit
	Id string `json:"id,omitempty"`
	// Metadata about upgrade unit
	Metadata []KeyValuePair `json:"metadata,omitempty"`
}
