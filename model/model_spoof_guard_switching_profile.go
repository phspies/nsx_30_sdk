package nsxt

// SpoofGuard configuration
type SpoofGuardSwitchingProfile struct {
	RequiredCapabilities []string `json:"required_capabilities,omitempty"`
	ResourceType string `json:"resource_type"`
	// List of providers for white listed address bindings.
	WhiteListProviders []string `json:"white_list_providers"`
}
