package nsxt

// Dropdown item definition
type DropdownItem struct {
	// An expression that represents the items of the dropdown filter.
	Field string `json:"field"`
	// expression to extract display name to be shown in the drop down.
	DisplayName string `json:"display_name,omitempty"`
	// Value of filter inside dropdown filter.
	Value string `json:"value"`
}
