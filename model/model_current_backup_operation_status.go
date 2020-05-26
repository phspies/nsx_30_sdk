package nsxt

// Current backup operation status
type CurrentBackupOperationStatus struct {
	// Current step of operation
	CurrentStep string `json:"current_step,omitempty"`
	// Unique identifier of current backup
	BackupId string `json:"backup_id,omitempty"`
	// Additional human-readable status information about current step
	CurrentStepMessage string `json:"current_step_message,omitempty"`
	// Time when operation is expected to end
	EndTime int64 `json:"end_time,omitempty"`
	// Type of operation that is in progress. Returns none if no operation is in progress, in which case none of the other fields will be set. 
	OperationType string `json:"operation_type"`
	// Time when operation was started
	StartTime int64 `json:"start_time,omitempty"`
}
