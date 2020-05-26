package nsxt

// Audit log in RFC5424 format
type AuditLog struct {
	// Priority field of the log
	Priority int64 `json:"priority"`
	// Process ID field of the log
	Procid int64 `json:"procid"`
	// Facility field of the log
	Facility int64 `json:"facility"`
	// Full log with both header and message
	FullLog string `json:"full_log"`
	// Message ID field of the log
	Msgid string `json:"msgid"`
	// Application name field of the log
	Appname string `json:"appname"`
	// Date and time in UTC of the log
	Timestamp string `json:"timestamp"`
	// Message field of the log
	Message string `json:"message"`
	// Hostname field of the log
	Hostname string `json:"hostname"`
	StructData *StructuredData `json:"struct_data"`
}
