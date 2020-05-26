package nsxt

// Static filters
type StaticFilter struct {
	// display name to be shown in the drop down for static filter.
	DisplayName string `json:"display_name,omitempty"`
	// Value of static filter inside dropdown filter.
	Value string `json:"value,omitempty"`
}
