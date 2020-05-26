package nsxt

// Inter SR IBGP configuration
type InterSrRoutingConfig struct {
	// While creation of BGP config this flag will be set to - true for Tier0 logical router with Active-Active high-availability mode - false for Tier0 logical router with Active-Standby high-availability mode. User can change this value while updating inter-sr config. 
	Enabled bool `json:"enabled,omitempty"`
}
