package nsxt

// Errors encountered while fetching entities in the forwarding path
type PortConnectionError struct {
	ErrorSummary string `json:"error_summary,omitempty"`
	ErrorDetails *interface{} `json:"error_details,omitempty"`
	EntityType string `json:"entity_type,omitempty"`
}
