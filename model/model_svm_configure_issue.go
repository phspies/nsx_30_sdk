package nsxt

// Type of issue and detailed description of the issues in case of post-VM  deployment configuration failure.
type SvmConfigureIssue struct {
	// List of errors along with details like errorId and error messages.
	Errors []SiErrorClass `json:"errors,omitempty"`
	// The ID of service instance which was deployed.
	ServiceInstanceId string `json:"service_instance_id,omitempty"`
}
