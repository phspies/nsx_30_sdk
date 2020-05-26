package nsxt

type Alarm struct {
	// Indicates system owned resource
	SystemOwned bool `json:"_system_owned,omitempty"`
	// Defaults to ID if not set
	DisplayName string `json:"display_name,omitempty"`
	// Detailed description of Alarm. This is the same detailed description as the corresponding Event identified by feature_name.event_type. 
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
	// ID that uniquely identifies an Alarm. 
	Id string `json:"id,omitempty"`
	// The type of this resource.
	ResourceType string `json:"resource_type,omitempty"`
	// Indicates when the corresponding Event instance was last reported in milliseconds since epoch. 
	LastReportedTime int64 `json:"last_reported_time,omitempty"`
	// Indicate the status which the Alarm is in. 
	Status string `json:"status"`
	// The entity that the Event instance applies to. Note entity_id may not be included in a response body. For example, the cpu_high Event may not return an entity_id. 
	EntityId string `json:"entity_id,omitempty"`
	// Name of Event, e.g. manager_cpu_usage_high, certificate_expired. 
	EventType string `json:"event_type,omitempty"`
	// Recommended action for Alarm. This is the same action as the corresponding Event identified by feature_name.event_type. 
	RecommendedAction string `json:"recommended_action,omitempty"`
	// Severity of the Alarm.Can be one of - CRITICAL, HIGH, MEDIUM, LOW. 
	Severity string `json:"severity,omitempty"`
	// The UUID of the node that the Event instance applies to. 
	NodeId string `json:"node_id,omitempty"`
	// Feature defining this Event, e.g. manager_health, certificates. 
	FeatureName string `json:"feature_name,omitempty"`
	// User ID of the user that set the status value to RESOLVED. This value can be SYSTEM to indicate that the system resolved the Alarm, for example when the system determines CPU usage is no longer high and the cpu_high Alarm is no longer applicable. This property is only returned when the status value is RESOLVED. 
	ResolvedBy string `json:"resolved_by,omitempty"`
	// Display name of Event type. 
	EventTypeDisplayName string `json:"event_type_display_name,omitempty"`
	// Display name of node that the event instance applies to. 
	NodeDisplayName string `json:"node_display_name,omitempty"`
	// Summary description of Alarm. This is the same summary description as the corresponding Event identified by feature_name.event_type. 
	Summary string `json:"summary,omitempty"`
	// Type of alarm source of the Event instance. Can be one of - INTENT_PATH, ENTITY_ID. 
	AlarmSourceType string `json:"alarm_source_type,omitempty"`
	// The resource type of node that the Event instance applies to eg. ClusterNodeConfig, HostNode, EdgeNode. 
	NodeResourceType string `json:"node_resource_type,omitempty"`
	// If alarm_source_type = INTENT_PATH, this field will contain a list of intent paths for the entity that the event instance applies to. If alarm_source_type = ENTITY_ID, this field will contain a list with a single item identifying the entity id that the event instance applies to. 
	AlarmSource []string `json:"alarm_source,omitempty"`
	// Display name of feature defining this Event. 
	FeatureDisplayName string `json:"feature_display_name,omitempty"`
	// User ID of the user that set the status value to SUPPRESSED. This property is only returned when the status value is SUPPRESSED. 
	SuppressedBy string `json:"suppressed_by,omitempty"`
	// Indicates when the Alarm was suppressed in milliseconds since epoch. This property is only returned when the status value is SUPPRESSED. 
	SuppressStartTime int64 `json:"suppress_start_time,omitempty"`
	// Indicates when the Alarm was resolved in milliseconds since epoch. This property is only returned when the status value is RESOLVED. 
	ResolvedTime int64 `json:"resolved_time,omitempty"`
	// The entity type that the Event instance applies to. 
	EntityResourceType string `json:"entity_resource_type,omitempty"`
	// The time period between suppress_start_time and suppress_start_time + suppress_duration (specified in hours) an Alarm is SUPPRESSED. This property is only returned when the status value is SUPPRESSED. 
	SuppressDuration int64 `json:"suppress_duration,omitempty"`
	// IP addresses of node that the event instance applies to. 
	NodeIpAddresses []string `json:"node_ip_addresses,omitempty"`
	// The number of reoccurrences since this alarm has been SUPPRESSED. 
	ReoccurrencesWhileSuppressed int64 `json:"reoccurrences_while_suppressed,omitempty"`
}
