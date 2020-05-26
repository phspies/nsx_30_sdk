package nsxt

// Details about switch to be migrated
type MigrationSwitchInfo struct {
	// Kind of switch, can be DVS, VSS.
	Kind string `json:"kind,omitempty"`
	// Version of the switch to be migrated.
	Version string `json:"version,omitempty"`
	// Number of PNICs associated with this switch.
	PnicCount int32 `json:"pnic_count,omitempty"`
	// Switch Identifier.
	Id string `json:"id,omitempty"`
	// Name of the switch.
	Name string `json:"name,omitempty"`
}
