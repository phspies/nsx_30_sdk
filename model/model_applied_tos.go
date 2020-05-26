package nsxt

// Entity lists where the profile will be enabled on. 
type AppliedTos struct {
	// Logical Port List
	LogicalPorts []ResourceReference `json:"logical_ports,omitempty"`
	// Logical Switch List
	LogicalSwitches []ResourceReference `json:"logical_switches,omitempty"`
	// NSGroup List
	Nsgroups []ResourceReference `json:"nsgroups,omitempty"`
}
