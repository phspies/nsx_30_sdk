package nsxt

type SwitchingProfileTypeIdEntry struct {
	// key value
	Value string `json:"value"`
	// Supported switching profiles. 'PortMirroringSwitchingProfile' is deprecated, please turn to \"Troubleshooting And Monitoring: Portmirroring\" and use PortMirroringSession API for port mirror function. 
	Key string `json:"key,omitempty"`
}
