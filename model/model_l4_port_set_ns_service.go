package nsxt

// L4PortSet can be specified in comma separated notation of parts. Parts of a L4PortSet includes single integer or range of port in hyphen notation. Example of a PortSet: \"22, 33-70, 44\". 
type L4PortSetNsService struct {
	// The specific type of NSServiceElement
	ResourceType string `json:"resource_type"`
	// Destination ports
	DestinationPorts []string `json:"destination_ports,omitempty"`
	L4Protocol string `json:"l4_protocol"`
	// Source ports
	SourcePorts []string `json:"source_ports,omitempty"`
}
