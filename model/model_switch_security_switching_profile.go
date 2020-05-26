package nsxt

type SwitchSecuritySwitchingProfile struct {
	RequiredCapabilities []string `json:"required_capabilities,omitempty"`
	ResourceType string `json:"resource_type"`
	BpduFilter *BpduFilter `json:"bpdu_filter,omitempty"`
	RateLimits *RateLimits `json:"rate_limits,omitempty"`
	// RA Guard when enabled blocks unauthorized/rogue Router Advertisement (RA) packets.
	RaGuardEnabled bool `json:"ra_guard_enabled,omitempty"`
	DhcpFilter *DhcpFilter `json:"dhcp_filter,omitempty"`
	// A flag to block all traffic except IP/(G)ARP/BPDU
	BlockNonIpTraffic bool `json:"block_non_ip_traffic,omitempty"`
}
