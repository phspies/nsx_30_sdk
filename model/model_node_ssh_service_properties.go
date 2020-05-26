package nsxt

// Node SSH service properties
type NodeSshServiceProperties struct {
	// Service name
	ServiceName string `json:"service_name"`
	ServiceProperties *SshServiceProperties `json:"service_properties,omitempty"`
}
