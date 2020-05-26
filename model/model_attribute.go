package nsxt

// Attribute specific to a partner. There attributes are passed on to the partner appliance and is opaque to the NSX Manager. The Attributes used by the partner applicance.
type Attribute struct {
	// Read only Attribute cannot be overdidden by service instance/deployment.
	ReadOnly bool `json:"read_only,omitempty"`
	// Attribute Type can be of any of the allowed enum type.
	AttributeType string `json:"attribute_type,omitempty"`
	// Attribute display name string value.
	DisplayName string `json:"display_name,omitempty"`
	// Attribute value string value.
	Value string `json:"value,omitempty"`
	// Attribute key string value.
	Key string `json:"key"`
}
