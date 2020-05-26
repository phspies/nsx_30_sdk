package nsxt

// Type of issue and detailed description of the issue in case of deployment failure.
type ServiceDeploymentIssue struct {
	// Description of issue encountered while service deployment.
	IssueDescription string `json:"issue_description,omitempty"`
	// Timestamp when issue was issue encountered while service deployment.
	IssueTimestamp string `json:"issue_timestamp,omitempty"`
	// Type of issue encountered while service deployment.
	IssueType string `json:"issue_type"`
}