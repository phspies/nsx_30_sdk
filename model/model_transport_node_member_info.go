package nsxt

// Information about participating transport nodes
type TransportNodeMemberInfo struct {
	// List of host switches using the transport zone
	HostSwitches []HostSwitchInfo `json:"host_switches,omitempty"`
	// Display name of the transport node which has one or more host switches which belong to associated transport zone.
	TransportNodeDisplayName string `json:"transport_node_display_name,omitempty"`
	// Id of the compute collection to which this transport node belongs. Empty if this is standalone transport node or non ESX type node.
	ComputeCollectionId string `json:"compute_collection_id,omitempty"`
	// Id of the transport node which has one or more host switches which belong to associated transport zone.
	TransportNodeId string `json:"transport_node_id,omitempty"`
}
