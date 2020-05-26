package nsxt

// Represents legend that describes the entities of the widget.
type Legend struct {
	// Describes the relative placement of legend. The legend of a widget can be placed either to the TOP or BOTTOM or LEFT or RIGHT relative to the widget. For example, if RIGHT is chosen then legend is placed to the right of the widget.
	Position string `json:"position,omitempty"`
	// If set to true, it will display the counts in legend. If set to false, counts of entities are not displayed in the legend.
	DisplayCount bool `json:"display_count,omitempty"`
	// Describes the render type for the legend. The legend for an entity describes the entity in the widget. The supported legend type is a circle against which the entity's details such as display_name are shown. The color of the circle denotes the color of the entity shown inside the widget.
	Type_ string `json:"type,omitempty"`
	// Show unit of entities in the legend.
	Unit string `json:"unit,omitempty"`
	// Describes the alignment of legend. Alignment of a legend denotes how individual items of the legend are aligned in a container. For example, if VERTICAL is chosen then the items of the legend will appear one below the other and if HORIZONTAL is chosen then the items will appear side by side.
	Alignment string `json:"alignment,omitempty"`
}
