package nsxt

type LogicalSwitchStatus struct {
	// Count of Logical Ports belonging to this switch
	NumLogicalPorts int32 `json:"num_logical_ports,omitempty"`
	// Unique ID identifying the the Logical Switch
	LogicalSwitchId string `json:"logical_switch_id,omitempty"`
}
