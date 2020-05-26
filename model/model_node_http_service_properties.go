package nsxt

// Node HTTP service properties
type NodeHttpServiceProperties struct {
	// Service name
	ServiceName string `json:"service_name"`
	ServiceProperties *HttpServiceProperties `json:"service_properties,omitempty"`
}
