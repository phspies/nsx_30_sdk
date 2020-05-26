package nsxt

type UpgradeHistory struct {
	// Timestamp (in milliseconds since epoch) when the upgrade was performed
	Timestamp int64 `json:"timestamp"`
	// Version being upgraded to
	TargetVersion string `json:"target_version"`
	// Version before the upgrade started
	InitialVersion string `json:"initial_version"`
	// Status of the upgrade
	UpgradeStatus string `json:"upgrade_status"`
}
