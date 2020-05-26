package nsxt

type UpgradeUnit struct {
	Self *SelfResourceLink `json:"_self,omitempty"`
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink `json:"_links,omitempty"`
	// Schema for this resource
	Schema string `json:"_schema,omitempty"`
	Group *UpgradeUnitGroupInfo `json:"group,omitempty"`
	// List of warnings indicating issues with the upgrade unit that may result in upgrade failure
	Warnings []string `json:"warnings,omitempty"`
	// This is component version e.g. if upgrade unit is of type edge, then this is edge version.
	CurrentVersion string `json:"current_version,omitempty"`
	// Metadata about upgrade unit
	Metadata []KeyValuePair `json:"metadata,omitempty"`
	// Upgrade unit type
	Type_ string `json:"type,omitempty"`
	// Identifier of the upgrade unit
	Id string `json:"id,omitempty"`
	// Name of the upgrade unit
	DisplayName string `json:"display_name,omitempty"`
}
