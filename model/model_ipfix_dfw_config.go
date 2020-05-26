package nsxt

// It defines IPFIX DFW Configuration.
type IpfixDfwConfig struct {
	// List of objects where the IPFIX Config will be enabled.
	AppliedTos []ResourceReference `json:"applied_tos,omitempty"`
	// Supported IPFIX Config Types.
	ResourceType string `json:"resource_type"`
	// This priority field is used to resolve conflicts in Logical Ports which are covered by more than one IPFIX profiles. The IPFIX exporter will send records to Collectors in highest priority profile (lowest number) only. 
	Priority int64 `json:"priority"`
	// Each IPFIX DFW config can have its own collector config. 
	Collector string `json:"collector"`
	// For long standing active flows, IPFIX records will be sent per timeout period 
	ActiveFlowExportTimeout int64 `json:"active_flow_export_timeout,omitempty"`
	TemplateParameters *IpfixDfwTemplateParameters `json:"template_parameters,omitempty"`
	// An identifier that is unique to the exporting process and used to meter the Flows. 
	ObservationDomainId int64 `json:"observation_domain_id"`
}
