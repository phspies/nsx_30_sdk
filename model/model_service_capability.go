package nsxt

// Service capabilities that will be inherited by service VMs created using a service definition that contains this service capability.
type ServiceCapability struct {
	// Indicating whether service supports NSH liveness detection.
	NshLivenessSupportEnabled bool `json:"nsh_liveness_support_enabled,omitempty"`
	// Indicating whether service is configured to decrement SI field in NSH metadata.
	CanDecrementSi bool `json:"can_decrement_si,omitempty"`
}
