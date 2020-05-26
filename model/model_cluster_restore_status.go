package nsxt

// Cluster restore status
type ClusterRestoreStatus struct {
	Status *GlobalRestoreStatus `json:"status,omitempty"`
	Step *RestoreStep `json:"step,omitempty"`
	// The list of allowed endpoints, based on the current state of the restore process 
	Endpoints []ResourceLink `json:"endpoints,omitempty"`
	// Total number of steps in the entire restore process
	TotalSteps int64 `json:"total_steps,omitempty"`
	// Timestamp when restore was started in epoch millisecond
	RestoreStartTime int64 `json:"restore_start_time,omitempty"`
	// Timestamp when restore was completed in epoch millisecond
	RestoreEndTime int64 `json:"restore_end_time,omitempty"`
	// Timestamp when backup was initiated in epoch millisecond
	BackupTimestamp int64 `json:"backup_timestamp,omitempty"`
	// Unique id for backup request
	Id string `json:"id,omitempty"`
	// Instructions for users to reconcile Restore operations
	Instructions []InstructionInfo `json:"instructions,omitempty"`
}