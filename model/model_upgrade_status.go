package nsxt

type UpgradeStatus struct {
	HostStatus *HostUpgradeStatus `json:"host_status,omitempty"`
	CcpStatus *CcpUpgradeStatus `json:"ccp_status,omitempty"`
	EdgeStatus *EdgeUpgradeStatus `json:"edge_status,omitempty"`
	// Status of upgrade
	OverallUpgradeStatus string `json:"overall_upgrade_status,omitempty"`
	// List of component statuses
	ComponentStatus []ComponentUpgradeStatus `json:"component_status,omitempty"`
}
