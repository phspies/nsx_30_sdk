package nsxt

// BPDU filter configuration
type BpduFilter struct {
	// Pre-defined list of allowed MAC addresses to be excluded from BPDU filtering
	WhiteList []string `json:"white_list,omitempty"`
	// Indicates whether BPDU filter is enabled
	Enabled bool `json:"enabled"`
}
