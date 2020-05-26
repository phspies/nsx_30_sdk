package nsxt

// Interface PIM[Protocol Independent Multicast] configuration parameters. 
type InterfacePimConfig struct {
	// If the flag is set to true - it will enable PIM on the uplink interface. If the flag is set to false - it will disable PIM on the uplink interface. 
	Enabled bool `json:"enabled,omitempty"`
}
