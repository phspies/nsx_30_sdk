package nsxt

// Port Connection Bare Metal Entities
type PortConnectionBmEntities struct {
	SrcPort *LogicalPort `json:"src_port,omitempty"`
	DstPort *LogicalPort `json:"dst_port,omitempty"`
}
