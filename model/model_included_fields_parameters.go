package nsxt

// A list of fields to include in query results
type IncludedFieldsParameters struct {
	// Comma separated list of fields that should be included in query result
	IncludedFields string `json:"included_fields,omitempty"`
}
