package nsxt

// Port Connection Logical Switch Entity
type PortConnectionLogicalSwitch struct {
	Resource *ManagedResource `json:"resource,omitempty"`
	// Resource ID is mapped to this. (ID is Generated for Edge node groups, since resource will be null)
	Id string `json:"id,omitempty"`
	// States of Logical Ports that are attached to a VIF/VM
	VmPortsStates []LogicalPortState `json:"vm_ports_states,omitempty"`
	// Logical Ports that are attached to a VIF/VM
	VmPorts []LogicalPort `json:"vm_ports,omitempty"`
	// Virutal Network Interfaces that are attached to the Logical Ports
	VmVnics []VirtualNetworkInterface `json:"vm_vnics,omitempty"`
	// Logical Ports that are attached to a router
	RouterPorts []LogicalPort `json:"router_ports,omitempty"`
}
