package nsxt

// Collection of edge nodes backing a logical router
type PortConnectionEdgeNodeGroup struct {
	Resource *ManagedResource `json:"resource,omitempty"`
	// Resource ID is mapped to this. (ID is Generated for Edge node groups, since resource will be null)
	Id string `json:"id,omitempty"`
	EdgeNodes []TransportNode `json:"edge_nodes,omitempty"`
	// Id of the logical router
	LogicalRouterId string `json:"logical_router_id"`
}
