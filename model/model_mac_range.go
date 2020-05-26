package nsxt

// A range of MAC addresses with a start and end value
type MacRange struct {
	// Start value for MAC address range
	Start string `json:"start"`
	// End value for MAC address range
	End string `json:"end"`
}
