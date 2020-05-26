package nsxt

// Render configuration to be applied to the widget.
type RenderConfiguration struct {
	// The color to use when rendering an entity. For example, set color as 'RED' to render a portion of donut in red.
	Color string `json:"color,omitempty"`
	// If the condition is met then the rendering specified for the condition will be applied. Examples of expression syntax are provided under 'example_request' section of 'CreateWidgetConfiguration' API.
	Condition string `json:"condition,omitempty"`
	// If specified, overrides the field value. This can be used to display a meaningful value in situations where field value is not available or not configured.
	DisplayValue string `json:"display_value,omitempty"`
	// Multi-line text to be shown on tooltip while hovering over the UI element if the condition is met.
	Tooltip []Tooltip `json:"tooltip,omitempty"`
	// Icons to be applied at dashboard for widgets and UI elements.
	Icons []Icon `json:"icons,omitempty"`
}
