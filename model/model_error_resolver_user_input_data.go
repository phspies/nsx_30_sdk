package nsxt

// Corresponds to one property entered by the user
type ErrorResolverUserInputData struct {
	// The datatype of the given property. Useful for data validation
	DataType string `json:"data_type"`
	// Name of the property supplied by the user
	PropertyName string `json:"property_name"`
	// The value associated with the above property
	PropertyValue string `json:"property_value,omitempty"`
}
