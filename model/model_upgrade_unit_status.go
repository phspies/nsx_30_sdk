package nsxt

type UpgradeUnitStatus struct {
	// Status of upgrade unit
	Status string `json:"status,omitempty"`
	// List of errors occurred during upgrade of this upgrade unit
	Errors []string `json:"errors,omitempty"`
	// Name of upgrade unit
	DisplayName string `json:"display_name,omitempty"`
	// Indicator of upgrade progress in percentage
	PercentComplete float64 `json:"percent_complete,omitempty"`
	// Identifier of upgrade unit
	Id string `json:"id,omitempty"`
	// Metadata about upgrade unit
	Metadata []KeyValuePair `json:"metadata,omitempty"`
}
