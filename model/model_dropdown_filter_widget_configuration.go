package nsxt

// Represents configuration for dropdown filter widget.
type DropdownFilterWidgetConfiguration struct {
	// Alias to be used when emitting filter value.
	Alias string `json:"alias,omitempty"`
	// Expression to specify default value of filter.
	DefaultValue string `json:"default_value,omitempty"`
	// Additional static items to be added in dropdown filter. Example can be 'ALL'.
	StaticFilters []StaticFilter `json:"static_filters,omitempty"`
	DropdownItem *DropdownItem `json:"dropdown_item"`
	// If the condition is met then the static filter will be added. If no condition is provided, then the static filters will be applied unconditionally.
	StaticFilterCondition string `json:"static_filter_condition,omitempty"`
	// Placeholder message to be displayed in dropdown filter.
	PlaceholderMsg string `json:"placeholder_msg,omitempty"`
}
