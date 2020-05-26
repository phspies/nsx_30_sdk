package nsxt

type L2ForwarderStatusPerNode struct {
	TransportNode *ResourceReference `json:"transport_node,omitempty"`
	// High Availability status of a service router on the edge node. 
	HighAvailabilityStatus string `json:"high_availability_status,omitempty"`
}
