package nsxt

// install-upgrade service properties
type InstallUpgradeServiceProperties struct {
	// IP of manager on which install-upgrade is enabled
	EnabledOn string `json:"enabled_on,omitempty"`
	// True if service enabled; otherwise, false
	Enabled bool `json:"enabled"`
}
