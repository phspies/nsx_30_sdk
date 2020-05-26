package nsxt

// The configuration for Internet protocol flow information export (IPFIX) and would be enabled on the switching module to collect flow information. 
type IpfixSwitchUpmProfile struct {
	// All IPFIX profile types.
	ResourceType string `json:"resource_type"`
	// This priority field is used to resolve conflicts in logical ports/switch  which inherit multiple switch IPFIX profiles from NSGroups.  Override rule is : for multiple profiles inherited from NSGroups, the one with highest priority (lowest number) overrides others; the profile directly applied to logical switch overrides profiles inherited from NSGroup; the profile directly applied to logical port overides profiles inherited from logical switch and/or nsgroup;  The IPFIX exporter will send records to collectors of final effective profile only. 
	Priority int32 `json:"priority"`
	// Each IPFIX switching profile can have its own collector profile. 
	CollectorProfile string `json:"collector_profile"`
	// The time in seconds after a flow is expired if no more packets matching this flow are received by the cache. 
	IdleTimeout int32 `json:"idle_timeout,omitempty"`
	// The maximum number of flow entries in each exporter flow cache. 
	MaxFlows int64 `json:"max_flows,omitempty"`
	// An identifier that is unique to the exporting process and used to meter the Flows. 
	ObservationDomainId int64 `json:"observation_domain_id"`
	// The time in seconds after a flow is expired even if more packets matching this Flow are received by the cache. 
	ActiveTimeout int32 `json:"active_timeout,omitempty"`
	// It controls whether sample result includes overlay flow info. 
	ExportOverlayFlow bool `json:"export_overlay_flow,omitempty"`
	AppliedTos *AppliedTos `json:"applied_tos,omitempty"`
	// The probability in percentage that a packet is sampled. The value should be  in range (0,100] and can only have three decimal places at most. The probability  is equal for every packet. 
	PacketSampleProbability float64 `json:"packet_sample_probability,omitempty"`
}
