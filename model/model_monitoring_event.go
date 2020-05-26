package nsxt

type MonitoringEvent struct {
	// Indicates system owned resource
	SystemOwned bool `json:"_system_owned,omitempty"`
	// Defaults to ID if not set
	DisplayName string `json:"display_name,omitempty"`
	// Detailed description of the event. 
	Description string `json:"description,omitempty"`
	// Opaque identifiers meaningful to the API user
	Tags []Tag `json:"tags,omitempty"`
	// ID of the user who created this resource
	CreateUser string `json:"_create_user,omitempty"`
	// Protection status is one of the following: PROTECTED - the client who retrieved the entity is not allowed             to modify it. NOT_PROTECTED - the client who retrieved the entity is allowed                 to modify it REQUIRE_OVERRIDE - the client who retrieved the entity is a super                    user and can modify it, but only when providing                    the request header X-Allow-Overwrite=true. UNKNOWN - the _protection field could not be determined for this           entity. 
	Protection string `json:"_protection,omitempty"`
	// Timestamp of resource creation
	CreateTime int64 `json:"_create_time,omitempty"`
	// Timestamp of last modification
	LastModifiedTime int64 `json:"_last_modified_time,omitempty"`
	// ID of the user who last modified this resource
	LastModifiedUser string `json:"_last_modified_user,omitempty"`
	// Unique identifier in the form of feature_name.event_type. 
	Id string `json:"id,omitempty"`
	// The type of this resource.
	ResourceType string `json:"resource_type,omitempty"`
	// Indicates if the threshold property is configurable via the API. 
	IsThresholdFixed bool `json:"is_threshold_fixed,omitempty"`
	// Optional field containing OID for SNMP trap sent when Event instance is False. This value is null if suppress_snmp_trap or suppress_clear_oid is True. 
	EventFalseSnmpOid string `json:"event_false_snmp_oid,omitempty"`
	// Description of Event when an Event instance transitions from True to False. 
	DescriptionOnClear string `json:"description_on_clear,omitempty"`
	// Name of Event, e.g. manager_cpu_usage_high, certificate_expired. 
	EventType string `json:"event_type,omitempty"`
	// Recommended action for the alarm condition. 
	RecommendedAction string `json:"recommended_action,omitempty"`
	// Severity of the Event.Can be one of - CRITICAL, HIGH, MEDIUM, LOW. 
	Severity string `json:"severity,omitempty"`
	// Percentage of samples to consider and used in combination with threshold when determining whether an Event instance status is True or False. Event evaluation uses sampling to determine Event instance status. A higher sensitivity value specifies that more samples are used to ensure accuracy and ignore infrequent or rare spikes in sampled data. 
	Sensitivity int64 `json:"sensitivity"`
	// Flag to indicate whether sampling for this Event is off or on. 
	IsDisabled bool `json:"is_disabled,omitempty"`
	// Flag to suppress Alarm generation. Alarms are not generated for this Event when this is set to True. 
	SuppressAlarm bool `json:"suppress_alarm,omitempty"`
	// Summary description of the event. 
	Summary string `json:"summary,omitempty"`
	// Display name of feature defining this Event. 
	FeatureDisplayName string `json:"feature_display_name,omitempty"`
	// Resource Type of entity where this event is applicable eg. LogicalSwitch, LogicalPort etc. 
	EntityResourceType string `json:"entity_resource_type,omitempty"`
	// Feature defining this Event, e.g. manager_health, certificates. 
	FeatureName string `json:"feature_name,omitempty"`
	// Threshold to determine if a single sample is True. For example, if the configured threshold is 95% and the current CPU sample is 99%, then the current sample is considered True. 
	Threshold int64 `json:"threshold"`
	// Display name of Event type. 
	EventTypeDisplayName string `json:"event_type_display_name,omitempty"`
	// Flag to suppress SNMP trap generation. SNMP traps are not sent for this Event when this is set to True. 
	SuppressSnmpTrap bool `json:"suppress_snmp_trap,omitempty"`
	// Optional field containing OID for SNMP trap sent when Event instance is True. This value is null if suppress_snmp_trap is True. 
	EventTrueSnmpOid string `json:"event_true_snmp_oid,omitempty"`
	// Array identifying the nodes on which this Event is applicable. Can be one or more of the following values - nsx_public_cloud_gateway, nsx_edge, nsx_esx, nsx_kvm, nsx_manager. 
	NodeTypes []string `json:"node_types,omitempty"`
}
