package nsxt

// Dscp value is ignored in case of 'TRUSTED' DscpMode.
type Dscp struct {
	// Internal Forwarding Priority
	Priority int32 `json:"priority,omitempty"`
	// Trust settings
	Mode string `json:"mode,omitempty"`
}
