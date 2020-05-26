package nsxt

// Enables traffic limit for incoming/outgoing broadcast and multicast packets. Use 0 to disable rate limiting for a specific traffic type
type RateLimits struct {
	// Incoming multicast traffic limit in packets per second
	RxMulticast int32 `json:"rx_multicast,omitempty"`
	// Outgoing multicast traffic limit in packets per second
	TxMulticast int32 `json:"tx_multicast,omitempty"`
	// Whether rate limiting is enabled
	Enabled bool `json:"enabled,omitempty"`
	// Outgoing broadcast traffic limit in packets per second
	TxBroadcast int32 `json:"tx_broadcast,omitempty"`
	// Incoming broadcast traffic limit in packets per second
	RxBroadcast int32 `json:"rx_broadcast,omitempty"`
}