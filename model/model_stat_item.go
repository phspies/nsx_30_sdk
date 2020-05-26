package nsxt

// Displayed as a single number. It can be used to show the characteristics of entities such as Logical Switches, Firewall Rules, and so on. For example, number of logical switches and their admin states.
type StatItem struct {
	// Id of drilldown widget, if any. Id should be a valid id of an existing widget.
	DrilldownId string `json:"drilldown_id,omitempty"`
	// If expression for total is specified, it evaluates it. Total can be omitted if not needed to be shown.
	Total string `json:"total,omitempty"`
	// Multi-line text to be shown on tooltip while hovering over the stat.
	Tooltip []Tooltip `json:"tooltip,omitempty"`
	// Expression for stat to be displayed.
	Value string `json:"value"`
}
