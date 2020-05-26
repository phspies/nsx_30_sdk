package nsxt

// Represents high level logical grouping of portions or segments of a donut / stats chart.
type DonutSection struct {
	// Field of the root of the api result set for forming parts.
	RowListField string `json:"row_list_field,omitempty"`
	// Array of portions or parts of the donut or stats chart.
	Parts []DonutPart `json:"parts"`
	// If true, the section will be appled as template for forming parts. Only one part will be formed from each element of 'row_list_field'.
	Template bool `json:"template,omitempty"`
}
