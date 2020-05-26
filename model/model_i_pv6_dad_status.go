package nsxt

// IPv6 DAD status
type IPv6DadStatus struct {
	// DAD status for IP address on the port. 
	Status string `json:"status,omitempty"`
	// Array of transport node id on which DAD status is reported for given IP address. 
	TransportNode []ResourceReference `json:"transport_node,omitempty"`
	// IP address on the port for which DAD status is reported. 
	IpAddress string `json:"ip_address,omitempty"`
}
