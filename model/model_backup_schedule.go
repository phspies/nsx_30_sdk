package nsxt

// Abstract base type for Weekly or Interval Backup Schedule
type BackupSchedule struct {
	// Schedule type
	ResourceType string `json:"resource_type"`
}
