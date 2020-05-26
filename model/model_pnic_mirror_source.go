package nsxt

type PnicMirrorSource struct {
	// Resource types of mirror source
	ResourceType string `json:"resource_type"`
	// Transport node identifier for the pnic located.
	NodeId string `json:"node_id"`
	// Whether to filter encapsulated packet.
	Encapsulated bool `json:"encapsulated"`
	// Source physical NIC device names
	SourcePnics []string `json:"source_pnics"`
}
