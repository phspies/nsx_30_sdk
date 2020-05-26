package nsxt

// Backup operation status
type BackupOperationStatus struct {
	// Unique identifier of a backup
	BackupId string `json:"backup_id"`
	// Time when operation was ended
	EndTime int64 `json:"end_time,omitempty"`
	// True if backup is successfully completed, else false
	Success bool `json:"success"`
	// Time when operation was started
	StartTime int64 `json:"start_time,omitempty"`
	// Error code details
	ErrorMessage string `json:"error_message,omitempty"`
	// Error code
	ErrorCode string `json:"error_code,omitempty"`
}
