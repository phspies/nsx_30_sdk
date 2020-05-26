package nsxt

// Defines a graph
type GraphDefinition struct {
	PointDefinition *PointDefinition `json:"point_definition"`
	Label *Label `json:"label,omitempty"`
	// Additional rendering or conditional evaluation of the field values to be performed, if any.
	RenderConfiguration []RenderConfiguration `json:"render_configuration,omitempty"`
}
