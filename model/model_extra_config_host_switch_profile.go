package nsxt

// Profile for extra configs in host switch
type ExtraConfigHostSwitchProfile struct {
	RequiredCapabilities []string `json:"required_capabilities,omitempty"`
	// Supported HostSwitch profiles.
	ResourceType string `json:"resource_type"`
	// list of extra configs
	ExtraConfigs []ExtraConfig `json:"extra_configs,omitempty"`
}
