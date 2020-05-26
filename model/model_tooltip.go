package nsxt

// Tooltip to be shown while hovering over the dashboard UI element.
type Tooltip struct {
	// Text to be shown on tooltip while hovering over UI element. The text would be wrapped if it exceeds 80 chars.
	Text string `json:"text"`
	// If true, displays tooltip text in bold
	Heading bool `json:"heading,omitempty"`
	// If the condition is met then the tooltip will be applied. If no condition is provided, then the tooltip will be applied unconditionally. Examples of expression syntax are provided under 'example_request' section of 'CreateWidgetConfiguration' API.
	Condition string `json:"condition,omitempty"`
}
