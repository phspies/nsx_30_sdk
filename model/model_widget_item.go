package nsxt

// Represents a reference to a widget that is held by a container or a multi-widget or a View.
type WidgetItem struct {
	// Aligns widget either left or right.
	Alignment string `json:"alignment,omitempty"`
	// If true, separates this widget in a container.
	Separator bool `json:"separator,omitempty"`
	// Id of the widget configuration that is held by a multi-widget or a container or a view.
	WidgetId string `json:"widget_id"`
	// Determines placement of widget or container relative to other widgets and containers. The lower the weight, the higher it is in the placement order.
	Weight int32 `json:"weight,omitempty"`
	Label *Label `json:"label,omitempty"`
}
