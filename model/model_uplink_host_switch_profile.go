package nsxt

// Profile for uplink policies
type UplinkHostSwitchProfile struct {
	RequiredCapabilities []string `json:"required_capabilities,omitempty"`
	// Supported HostSwitch profiles.
	ResourceType string `json:"resource_type"`
	// list of LACP group
	Lags []Lag `json:"lags,omitempty"`
	// VLAN used for tagging Overlay traffic of associated HostSwitch
	TransportVlan int64 `json:"transport_vlan,omitempty"`
	Teaming *TeamingPolicy `json:"teaming"`
	// The protocol used to encapsulate overlay traffic
	OverlayEncap string `json:"overlay_encap,omitempty"`
	// List of named uplink teaming policies that can be used by logical switches
	NamedTeamings []NamedTeamingPolicy `json:"named_teamings,omitempty"`
	// Maximum Transmission Unit used for uplinks
	Mtu int32 `json:"mtu,omitempty"`
}
