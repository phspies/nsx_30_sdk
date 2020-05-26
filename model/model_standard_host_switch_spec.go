package nsxt

// Standard host switch specification is used for NSX configured transport node.
type StandardHostSwitchSpec struct {
	ResourceType string `json:"resource_type"`
	// Transport Node host switches
	HostSwitches []StandardHostSwitch `json:"host_switches"`
}
