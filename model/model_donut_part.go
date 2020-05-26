package nsxt

// Represents an entity or portion to be plotted on a donut or stats chart.
type DonutPart struct {
	// A numerical value that represents the portion or entity of the donut or stats chart.
	Field string `json:"field"`
	// If true, legend will be shown only if the data for the part is available. This is applicable only if legends are specified in widget configuration.
	HideEmptyLegend bool `json:"hide_empty_legend,omitempty"`
	// If the condition is met then the part will be displayed. Examples of expression syntax are provided under 'example_request' section of 'CreateWidgetConfiguration' API.
	Condition string `json:"condition,omitempty"`
	// Id of drilldown widget, if any. Id should be a valid id of an existing widget. A widget is considered as drilldown widget when it is associated with any other widget and provides more detailed information about any data item from the parent widget.
	DrilldownId string `json:"drilldown_id,omitempty"`
	Label *Label `json:"label,omitempty"`
	// Hyperlink of the specified UI page that provides details. If drilldown_id is provided, then navigation cannot be used.
	Navigation string `json:"navigation,omitempty"`
	// Multi-line text to be shown on tooltip while hovering over the portion.
	Tooltip []Tooltip `json:"tooltip,omitempty"`
	// Additional rendering or conditional evaluation of the field values to be performed, if any.
	RenderConfiguration []RenderConfiguration `json:"render_configuration,omitempty"`
}
