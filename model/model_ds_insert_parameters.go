package nsxt

// Parameters to tell where rule/section need to be added. All the params take rule/section Id.
type DsInsertParameters struct {
	// Operation
	Operation string `json:"operation,omitempty"`
	// Identifier of the anchor rule or section. This is a required field in case operation like 'insert_before' and 'insert_after'.
	Id string `json:"id,omitempty"`
}
