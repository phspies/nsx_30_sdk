package nsxt

// Host Switch for LLDP
type LldpHostSwitchProfile struct {
	RequiredCapabilities []string `json:"required_capabilities,omitempty"`
	// Supported HostSwitch profiles.
	ResourceType string `json:"resource_type"`
	// Enabled or disabled sending LLDP packets
	SendEnabled bool `json:"send_enabled"`
}
