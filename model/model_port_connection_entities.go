package nsxt

// Port Connection Entities (to help draw a visual picture of entities between two ports)
type PortConnectionEntities struct {
	Containers *PortConnectionContainersEntities `json:"containers"`
	Hypervisors []PortConnectionHypervisor `json:"hypervisors"`
	Errors []PortConnectionError `json:"errors"`
	LogicalSwitches []PortConnectionLogicalSwitch `json:"logical_switches"`
	EdgeNodeGroups []PortConnectionEdgeNodeGroup `json:"edge_node_groups,omitempty"`
	Routers []PortConnectionRouter `json:"routers,omitempty"`
	Vms []VirtualMachine `json:"vms"`
	Tunnels []PortConnectionTunnel `json:"tunnels"`
	PhysicalHosts *PortConnectionBmEntities `json:"physical_hosts,omitempty"`
}
