package nsxt

// NSX global configs for Distributed Firewall
type FirewallGlobalConfig struct {
	// Valid Global configuration types
	ResourceType string `json:"resource_type"`
	// This property is deprecated. The fast path mode is always enabled in Distributed Firewall.
	GlobalFastpathModeEnabled bool `json:"global_fastpath_mode_enabled,omitempty"`
	// When this flag is set to true, global address set is enabled in Distributed Firewall.
	GlobalAddrsetModeEnabled bool `json:"global_addrset_mode_enabled,omitempty"`
}
