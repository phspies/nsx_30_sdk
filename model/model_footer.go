package nsxt

// Footer of a widget that provides additional information or allows an action such as clickable url for navigation. An example usage of footer is provided under 'example_request' section of 'CreateWidgetConfiguration' API.
type Footer struct {
	// If the condition is met then the footer will be applied. Examples of expression syntax are provided under 'example_request' section of 'CreateWidgetConfiguration' API.
	Condition string `json:"condition,omitempty"`
	// Action to be performed at the footer of a widget. An action at the footer can be simple text description or a hyperlink to a UI page. Action allows a clickable url for navigation. An example usage of footer action is provided under 'example_request' section of 'CreateWidgetConfiguration' API.
	Actions []FooterAction `json:"actions,omitempty"`
}
