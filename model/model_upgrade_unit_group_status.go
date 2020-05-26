package nsxt

type UpgradeUnitGroupStatus struct {
	// Upgrade status of upgrade unit group
	Status string `json:"status,omitempty"`
	// Number of upgrade units in the group
	UpgradeUnitCount int32 `json:"upgrade_unit_count,omitempty"`
	// Number of nodes in the upgrade unit group that failed upgrade
	FailedCount int32 `json:"failed_count,omitempty"`
	// Indicator of upgrade progress in percentage
	PercentComplete float64 `json:"percent_complete,omitempty"`
	// Identifier for upgrade unit group
	GroupId string `json:"group_id,omitempty"`
	// Name of the upgrade unit group
	GroupName string `json:"group_name,omitempty"`
}
