package nsxt

// Node NTP service properties
type NodeNtpServiceProperties struct {
	// Service name
	ServiceName string `json:"service_name"`
	ServiceProperties *NtpServiceProperties `json:"service_properties,omitempty"`
}
