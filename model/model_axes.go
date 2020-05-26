package nsxt

// Represents X and Y axes of a graph. For a multi-graph, the same axes are shared by all the graphs.
type Axes struct {
	XLabel *Label `json:"x_label,omitempty"`
	YLabel *Label `json:"y_label,omitempty"`
}
