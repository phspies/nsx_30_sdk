package nsxt

// Preconfigured host switch specification is used for manually configured transport node. It is user's responsibility to ensure correct configuration is provided to NSX. This type is only valid for supported KVM fabric nodes.
type PreconfiguredHostSwitchSpec struct {
	ResourceType string `json:"resource_type"`
	// Preconfigured Transport Node host switches
	HostSwitches []PreconfiguredHostSwitch `json:"host_switches"`
}
