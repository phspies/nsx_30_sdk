package nsxt

// Status of edge upgrade
type EdgeUpgradeStatus struct {
	// Upgrade status of component
	Status string `json:"status,omitempty"`
	PreUpgradeStatus *UpgradeChecksExecutionStatus `json:"pre_upgrade_status,omitempty"`
	// Details about the upgrade status
	Details string `json:"details,omitempty"`
	// Component type for the upgrade status
	ComponentType string `json:"component_type,omitempty"`
	// Number of nodes of the type and at the component version
	NodeCountAtTargetVersion int32 `json:"node_count_at_target_version,omitempty"`
	// Target component version
	TargetComponentVersion string `json:"target_component_version,omitempty"`
	// Indicator of upgrade progress in percentage
	PercentComplete float64 `json:"percent_complete,omitempty"`
	// Can the upgrade of the remaining units in this component be skipped
	CanSkip bool `json:"can_skip,omitempty"`
	CurrentVersionNodeSummary *NodeSummaryList `json:"current_version_node_summary,omitempty"`
}
