package nsxt

type PnicMirrorDestination struct {
	// Resource types of mirror destination
	ResourceType string `json:"resource_type"`
	// Physical NIC device names to which to send the mirrored packets
	DestPnics []string `json:"dest_pnics"`
	// Transport node to which to send the mirrored packets
	NodeId string `json:"node_id"`
}
