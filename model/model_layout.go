package nsxt

// Represents layout of a container or widget
type Layout struct {
	// Describes layout of a container or widget. Layout describes how individual widgets are placed inside the container. For example, if HORIZONTAL is chosen widgets are placed side by side inside the container. If VERTICAL is chosen then widgets are placed one below the other. If GRID is chosen then the container or widget display area is divided into a grid of m rows and n columns, as specified in the properties, and the widgets are placed inside the grid.
	Type_ string `json:"type,omitempty"`
	Properties *LayoutProperties `json:"properties,omitempty"`
}
