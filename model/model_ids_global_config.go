package nsxt

// NSX global configs for Distributed Intrusion Services
type IdsGlobalConfig struct {
	// Valid Global configuration types
	ResourceType string `json:"resource_type"`
	// When this flag is set to true, IDS events would be sent to syslog.
	GlobalIdseventsToSyslogEnabled bool `json:"global_idsevents_to_syslog_enabled,omitempty"`
}
