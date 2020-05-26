package nsxt

// Represents a column of the Grid
type ColumnItem struct {
	// Sorting on column is based on the sort_key. sort_key represents the field in the output data on which sort is requested.
	SortKey string `json:"sort_key,omitempty"`
	// Data type of the field.
	Type_ string `json:"type"`
	// Multi-line text to be shown on tooltip while hovering over a cell in the grid.
	Tooltip []Tooltip `json:"tooltip,omitempty"`
	Label *Label `json:"label"`
	// Field from which values of the column will be derived.
	Field string `json:"field"`
	// If true, the value of the column are sorted in ascending order. Otherwise, in descending order.
	SortAscending bool `json:"sort_ascending,omitempty"`
	// Id of drilldown widget, if any. Id should be a valid id of an existing widget.
	DrilldownId string `json:"drilldown_id,omitempty"`
	// If set to true, hides the column
	Hidden bool `json:"hidden,omitempty"`
	// Hyperlink of the specified UI page that provides details. If drilldown_id is provided, then navigation cannot be used.
	Navigation string `json:"navigation,omitempty"`
	// Identifies the column and used for fetching content upon an user click or drilldown. If column identifier is not provided, the column's data will not participate in searches and drilldowns.
	ColumnIdentifier string `json:"column_identifier,omitempty"`
	// Render configuration to be applied, if any.
	RenderConfiguration []RenderConfiguration `json:"render_configuration,omitempty"`
}
