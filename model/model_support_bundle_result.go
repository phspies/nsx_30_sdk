package nsxt

type SupportBundleResult struct {
	RequestProperties *SupportBundleRequest `json:"request_properties,omitempty"`
	// Nodes where bundles were not generated or not copied to remote server
	FailedNodes []FailedNodeSupportBundleResult `json:"failed_nodes,omitempty"`
	// Nodes whose bundles were successfully copied to remote file server
	SuccessNodes []SuccessNodeSupportBundleResult `json:"success_nodes,omitempty"`
	// Nodes where bundle generation is pending or in progress
	RemainingNodes []RemainingSupportBundleNode `json:"remaining_nodes,omitempty"`
}
