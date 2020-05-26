package nsxt

// The option is used to filter data on given node.
type PacketCaptureOption struct {
	// The avaiable option names in the enum can be used to filter the capture data.
	Name string `json:"name,omitempty"`
	// Define the capture value according to the given capture option.
	Value string `json:"value,omitempty"`
}
