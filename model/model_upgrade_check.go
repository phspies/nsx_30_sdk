package nsxt

// Check to identify potential pre/post-upgrade issues
type UpgradeCheck struct {
	// Status of pre/post-upgrade check
	Status string `json:"status,omitempty"`
	// List of failures
	Failures []UpgradeCheckFailureMessage `json:"failures,omitempty"`
	// Name of the pre/post-upgrade check
	DisplayName string `json:"display_name,omitempty"`
	// List of failure messages. This field is deprecated now. Please use failures instead.
	FailureMessages []string `json:"failure_messages,omitempty"`
	// Component type
	ComponentType string `json:"component_type"`
}
