package nsxt

// Overrides the router advertisement attributes for the IPv6 prefixes. 
type NdraPrefixConfig struct {
	// Override the neighbor discovery prefix preferred time and prefix valid time for the subnet on uplink port whose network matches with the network address of CIDR specified in network_prefix. 
	NetworkPrefix string `json:"network_prefix"`
	// The time interval in seconds, in which the prefix is advertised as valid. 
	PrefixValidTime int64 `json:"prefix_valid_time,omitempty"`
	// The time interval in seconds, in which the prefix is advertised as preferred. 
	PrefixPreferredTime int64 `json:"prefix_preferred_time,omitempty"`
}
