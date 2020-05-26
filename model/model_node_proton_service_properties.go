package nsxt

// Node service properties
type NodeProtonServiceProperties struct {
	// Service name
	ServiceName string `json:"service_name"`
	ServiceProperties *LoggingServiceProperties `json:"service_properties,omitempty"`
}
