package nsxt

// Action specified at the footer of a widget to provide additional information or to provide a clickable url for navigation. An example usage of footer action is provided under the 'example_request' section of 'CreateWidgetConfiguration' API.
type FooterAction struct {
	// Hyperlink to the UI page that provides details of action.
	Url string `json:"url,omitempty"`
	// If true, the footer will appear in the underlying container that holds the widget.
	DockToContainerFooter bool `json:"dock_to_container_footer,omitempty"`
	Label *Label `json:"label"`
}
