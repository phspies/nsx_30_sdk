package nsxt

// Schedule to specify the interval time at which automated backups need to be taken
type IntervalBackupSchedule struct {
	// Schedule type
	ResourceType string `json:"resource_type"`
	// Time interval in seconds between two consecutive automated backups
	SecondsBetweenBackups int64 `json:"seconds_between_backups,omitempty"`
}
