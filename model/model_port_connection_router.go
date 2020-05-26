package nsxt

// Port Connection Logical Router Entity
type PortConnectionRouter struct {
	Resource *ManagedResource `json:"resource,omitempty"`
	// Resource ID is mapped to this. (ID is Generated for Edge node groups, since resource will be null)
	Id string `json:"id,omitempty"`
	// Uplink ports of the Logical Router.
	UplinkPorts []LogicalRouterPort `json:"uplink_ports,omitempty"`
	// Downlink ports of the Logical Router.
	DownlinkPorts []LogicalRouterPort `json:"downlink_ports,omitempty"`
}
