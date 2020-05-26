package nsxt

// Node AAA provider vIDM status
type NodeAuthProviderVidmStatus struct {
	// AAA provider vIDM status
	RuntimeState string `json:"runtime_state"`
	// vIDM enable flag
	VidmEnable bool `json:"vidm_enable"`
}
