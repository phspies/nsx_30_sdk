package nsxt

type FailedNodeSupportBundleResult struct {
	// Display name of node
	NodeDisplayName string `json:"node_display_name,omitempty"`
	// UUID of node
	NodeId string `json:"node_id,omitempty"`
	// Error code
	ErrorCode string `json:"error_code,omitempty"`
	// Error message
	ErrorMessage string `json:"error_message,omitempty"`
}
