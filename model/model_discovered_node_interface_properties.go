package nsxt

// Network interface properties of discovered node
type DiscoveredNodeInterfaceProperties struct {
	// Mac address of the interface
	PhysicalAddress string `json:"physical_address,omitempty"`
	// Switch name which is connected to nic, switch can be opaque, proxyHostSwitch or virtual
	ConnectedSwitch string `json:"connected_switch,omitempty"`
	// Type of virtual switch can be VSS, DVS or N-VDS.
	ConnectedSwitchType string `json:"connected_switch_type,omitempty"`
	// Id of the network interface
	InterfaceId string `json:"interface_id,omitempty"`
}
