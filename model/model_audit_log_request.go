package nsxt

type AuditLogRequest struct {
	// Audit logs should meet the filter condition
	LogFilter string `json:"log_filter,omitempty"`
	// Include logs with timstamps not past the age limit in days
	LogAgeLimit int64 `json:"log_age_limit,omitempty"`
	// Type of log filter
	LogFilterType string `json:"log_filter_type,omitempty"`
}
