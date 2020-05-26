package nsxt

// Represents a label-value pair.
type PropertyItem struct {
	// Represents field value of the property.
	Field string `json:"field"`
	// If true, separates this property in a widget.
	Separator bool `json:"separator,omitempty"`
	// Hyperlink of the specified UI page that provides details. This will be linked with value of the property.
	Navigation string `json:"navigation,omitempty"`
	// Render configuration to be applied, if any.
	RenderConfiguration []RenderConfiguration `json:"render_configuration,omitempty"`
	// Data type of the field.
	Type_ string `json:"type"`
	// Set to true if the field is a heading. Default is false.
	Heading bool `json:"heading,omitempty"`
	// If the condition is met then the property will be displayed. Examples of expression syntax are provided under 'example_request' section of 'CreateWidgetConfiguration' API.
	Condition string `json:"condition,omitempty"`
	Label *Label `json:"label,omitempty"`
}
