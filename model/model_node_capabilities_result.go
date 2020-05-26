package nsxt

// List of capabilities of a fabric node
type NodeCapabilitiesResult struct {
	// Node capability results
	Capabilities []NodeCapability `json:"capabilities"`
}
