package nsxt

// Defines the point of a graph.
type PointDefinition struct {
	// Represents the variable for the Y value of points that are plotted on the graph.
	YValue string `json:"y_value"`
	// An expression that represents the points of the graph
	Field string `json:"field"`
	// Id of drilldown widget, if any. Id should be a valid id of an existing widget. A widget is considered as drilldown widget when it is associated with any other widget and provides more detailed information about any data item from the parent widget.
	DrilldownId string `json:"drilldown_id,omitempty"`
	// Represents the variable for the X value of points that are plotted on the graph.
	XValue string `json:"x_value"`
	// Hyperlink of the specified UI page that provides details.
	Navigation string `json:"navigation,omitempty"`
	// Multi-line text to be shown on tooltip while hovering over the point of a graph.
	Tooltip []Tooltip `json:"tooltip,omitempty"`
}
