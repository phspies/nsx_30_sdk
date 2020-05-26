package nsxt

// Tunnel information between two given transport nodes
type PortConnectionTunnel struct {
	// Id of the source transport node
	SrcNodeId string `json:"src_node_id"`
	TunnelProperties *TunnelProperties `json:"tunnel_properties"`
}
