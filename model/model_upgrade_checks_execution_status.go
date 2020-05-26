package nsxt

// Execution status of pre/post-upgrade checks
type UpgradeChecksExecutionStatus struct {
	// Status of execution of pre/post-upgrade checks
	Status string `json:"status,omitempty"`
	// Number of nodes which generated failures or warnings in last execution of pre/post-upgrade checks 
	NodeWithIssuesCount int32 `json:"node_with_issues_count,omitempty"`
	// Details about current execution of pre/post-upgrade checks
	Details string `json:"details,omitempty"`
	// Total count of generated failures or warnings in last execution of pre/post-upgrade checks 
	FailureCount int32 `json:"failure_count,omitempty"`
	// Time (in milliseconds since epoch) when the execution of pre/post-upgrade checks started 
	StartTime int64 `json:"start_time,omitempty"`
	// Time (in milliseconds since epoch) when the execution of pre/post-upgrade checks completed 
	EndTime int64 `json:"end_time,omitempty"`
}
