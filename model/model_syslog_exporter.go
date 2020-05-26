package nsxt

// This object specifies what, where and how logs on NSX nodes are forwarded.
type SyslogExporter struct {
	// Maximum logging level for messages to be exported.
	MaxLogLevel string `json:"max_log_level"`
	// Protocol to be used to export logs to syslog server.
	Protocol string `json:"protocol"`
	// Server port on which syslog listener is listening.
	Port int64 `json:"port,omitempty"`
	// Syslog server IP address or hostname.
	Server string `json:"server"`
}
