package nsxt

// Error while applying transport node profile on discovered node
type ValidationError struct {
	// Discovered Node Id
	DiscoveredNodeId string `json:"discovered_node_id,omitempty"`
	// Validation error message
	ErrorMessage string `json:"error_message,omitempty"`
}
