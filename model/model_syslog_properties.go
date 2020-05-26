package nsxt

// This object contains list of syslog exporters used by NSX nodes.
type SyslogProperties struct {
	// List of syslog exporters.
	Exporters []SyslogExporter `json:"exporters"`
}
