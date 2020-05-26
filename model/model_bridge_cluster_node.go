package nsxt

// Bridge transport node
type BridgeClusterNode struct {
	// MAC address used for HA protocol
	HaMac string `json:"ha_mac,omitempty"`
	// UUID of the transport node
	TransportNodeId string `json:"transport_node_id"`
}
