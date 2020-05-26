package nsxt

// All the network interfaces of the discovered node
type DiscoveredNodeNetworkInterfaces struct {
	// Id of the discovered node
	DiscoveredNodeId string `json:"discovered_node_id"`
	// Network interfaces of the node
	NetworkInterfaces []DiscoveredNodeInterfaceProperties `json:"network_interfaces,omitempty"`
}
