package nsxt

type PortMirroringSwitchingProfile struct {
	RequiredCapabilities []string `json:"required_capabilities,omitempty"`
	ResourceType string `json:"resource_type"`
	// If this property not set, original package will not be truncated.
	SnapLength int64 `json:"snap_length,omitempty"`
	// port mirroring direction
	Direction string `json:"direction,omitempty"`
	// User-configurable 32-bit key
	Key int64 `json:"key,omitempty"`
	// List of destination addresses
	Destinations []string `json:"destinations,omitempty"`
}
