package nsxt

// Icon to be applied at dashboard for widgets and UI elements.
type Icon struct {
	// If specified as PRE, the icon appears before the UI element. If set as POST, the icon appears after the UI element.
	Placement string `json:"placement,omitempty"`
	// Icon will be rendered based on its type. For example, if ERROR is chosen, then icon representing error will be rendered.
	Type_ string `json:"type,omitempty"`
	// Multi-line text to be shown on tooltip while hovering over the icon.
	Tooltip []Tooltip `json:"tooltip,omitempty"`
}