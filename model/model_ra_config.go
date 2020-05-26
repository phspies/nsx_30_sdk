package nsxt

type RaConfig struct {
	// The maximum number of hops through which packets can pass before being discarded. 
	HopLimit int64 `json:"hop_limit,omitempty"`
	// Router lifetime value in seconds. A value of 0 indicates the router is not a default router for the receiving end. Any other value in this field specifies the lifetime, in seconds, associated with this router as a default router. 
	RouterLifetime int64 `json:"router_lifetime,omitempty"`
	// Interval between 2 Router advertisement in seconds. 
	RaInterval int64 `json:"ra_interval,omitempty"`
	// The time interval in seconds, in which the prefix is advertised as preferred. 
	PrefixPreferredTime int64 `json:"prefix_preferred_time,omitempty"`
	// The time interval in seconds, in which the prefix is advertised as valid. 
	PrefixLifetime int64 `json:"prefix_lifetime,omitempty"`
}
