package nsxt

// Section Realization Status
type DsSectionState struct {
	// Gives details of state of desired configuration. Additional enums with more details on progress/success/error states are sent for edge node. The success states are NODE_READY and TRANSPORT_NODE_READY, pending states are {VM_DEPLOYMENT_QUEUED, VM_DEPLOYMENT_IN_PROGRESS, REGISTRATION_PENDING} and other values indicate failures. \"in_sync\" state indicates that the desired configuration has been received by the host to which it applies, but is not yet in effect. When the configuration is actually in effect, the state will change to \"success\". Please note, failed state is deprecated. 
	State string `json:"state,omitempty"`
	// Array of configuration state of various sub systems
	Details []ConfigurationStateElement `json:"details,omitempty"`
	// Error code
	FailureCode int64 `json:"failure_code,omitempty"`
	// Error message in case of failure
	FailureMessage string `json:"failure_message,omitempty"`
	// This attribute represents revision number of section's desired state.
	RevisionDesired int64 `json:"revision_desired,omitempty"`
}
