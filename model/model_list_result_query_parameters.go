package nsxt

// Parameters that affect how list results are processed
type ListResultQueryParameters struct {
	// Comma-separated field names to include in query result
	Fields string `json:"fields,omitempty"`
}
