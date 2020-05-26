package nsxt

// Filter request parameters
type FilterRequest struct {
	// Comma seperated fields to be filtered on
	FieldNames string `json:"field_names"`
	// Filter value
	Value string `json:"value"`
}
