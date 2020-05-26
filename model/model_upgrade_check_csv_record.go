package nsxt

// CSV record for a pre/post-upgrade check
type UpgradeCheckCsvRecord struct {
	// Status of the pre/post-upgrade check
	Status string `json:"status,omitempty"`
	// Description of the pre/post-upgrade check
	CheckDescription string `json:"check_description,omitempty"`
	// Space-separated list of failure messages
	FailureMessages string `json:"failure_messages,omitempty"`
	// Display name of the pre/post-upgrade check
	CheckName string `json:"check_name"`
	// Identifier of the upgrade unit
	UpgradeUnitId string `json:"upgrade_unit_id,omitempty"`
	// Meta-data of the upgrade-unit
	UpgradeUnitMetadata string `json:"upgrade_unit_metadata,omitempty"`
	// Component type of the upgrade unit
	UpgradeUnitType string `json:"upgrade_unit_type"`
}
