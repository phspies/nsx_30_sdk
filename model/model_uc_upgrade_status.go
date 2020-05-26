package nsxt

// Upgrade status of upgrade-coordinator
type UcUpgradeStatus struct {
	// Status of UC upgrade
	Status string `json:"status,omitempty"`
	// Current state of UC upgrade
	State string `json:"state,omitempty"`
}
