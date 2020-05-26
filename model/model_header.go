package nsxt

// Header of a widget that provides additional information. This will be shown at the container level. It includes details as label value pairs.
type Header struct {
	// Alignment of header labels.
	ContentAlignment string `json:"content_alignment,omitempty"`
	// If the condition is met then the header will be applied. Examples of expression syntax are provided under 'example_request' section of 'CreateWidgetConfiguration' API.
	Condition string `json:"condition,omitempty"`
	// An array of label-value properties.
	SubHeaders []PropertyItem `json:"sub_headers,omitempty"`
}
