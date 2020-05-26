package nsxt

// Node install-upgrade service properties
type NodeInstallUpgradeServiceProperties struct {
	// Service name
	ServiceName string `json:"service_name"`
	ServiceProperties *InstallUpgradeServiceProperties `json:"service_properties,omitempty"`
}
