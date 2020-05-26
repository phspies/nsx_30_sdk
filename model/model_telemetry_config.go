package nsxt

type TelemetryConfig struct {
	// Indicates system owned resource
	SystemOwned bool `json:"_system_owned,omitempty"`
	// Defaults to ID if not set
	DisplayName string `json:"display_name,omitempty"`
	// Description of this resource
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
	// Unique identifier of this resource
	Id string `json:"id,omitempty"`
	// The type of this resource.
	ResourceType string `json:"resource_type,omitempty"`
	// Enable this to schedule data collection and upload times. If enabled, and a schedule is not provided, a default schedule (WEEKLY, Sunday at 2:00 a.m) will be applied. 
	ScheduleEnabled bool `json:"schedule_enabled"`
	TelemetryProxy *TelemetryProxy `json:"telemetry_proxy,omitempty"`
	// Enable this flag to participate in the Customer Experience Improvement Program. 
	CeipAcceptance bool `json:"ceip_acceptance"`
	TelemetrySchedule *TelemetrySchedule `json:"telemetry_schedule,omitempty"`
	// Enable this flag to specify a proxy, and provide the proxy settings.
	ProxyEnabled bool `json:"proxy_enabled,omitempty"`
}
