package nsxt

type MacManagementSwitchingProfile struct {
	RequiredCapabilities []string `json:"required_capabilities,omitempty"`
	ResourceType string `json:"resource_type"`
	MacLearning *MacLearningSpec `json:"mac_learning,omitempty"`
	// Allowing source MAC address change
	MacChangeAllowed bool `json:"mac_change_allowed,omitempty"`
}
