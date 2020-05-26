package nsxt

// Profile for BFD HA cluster setting
type EdgeHighAvailabilityProfile struct {
	// Supported cluster profiles.
	ResourceType string `json:"resource_type"`
	StandbyRelocationConfig *StandbyRelocationConfig `json:"standby_relocation_config,omitempty"`
	// Number of times a packet is missed before BFD declares the neighbor down.
	BfdDeclareDeadMultiple int64 `json:"bfd_declare_dead_multiple,omitempty"`
	// the time interval (in millisec) between probe packets for heartbeat purpose
	BfdProbeInterval int64 `json:"bfd_probe_interval,omitempty"`
	// BFD allowed hops
	BfdAllowedHops int64 `json:"bfd_allowed_hops,omitempty"`
}
