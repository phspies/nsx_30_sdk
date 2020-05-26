package nsxt

// IpfixObsPointConfig (i.e. global switch IPFIX config) is deprecated. Please use IpfixSwitchUpmProfile & IpfixCollectorUpmProfile instead. With them, switch IPFIX profile can be applied to specific entities, such as logical switch, logical port and so on. 
type IpfixObsPointConfig struct {
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
	// The time in seconds after a Flow is expired if no more packets matching this Flow are received by the cache. 
	IdleTimeout int32 `json:"idle_timeout,omitempty"`
	// An identifier that is unique to the exporting process and used to meter the Flows. 
	ObservationDomainId int64 `json:"observation_domain_id,omitempty"`
	// List of IPFIX collectors
	Collectors []IpfixCollector `json:"collectors,omitempty"`
	// The time in seconds after a Flow is expired even if more packets matching this Flow are received by the cache. 
	ActiveTimeout int32 `json:"active_timeout,omitempty"`
	// The probability in percentage that a packet is sampled. The value should be  in range (0,100] and can only have three decimal places at most. The probability  is equal for every packet. 
	PacketSampleProbability float64 `json:"packet_sample_probability,omitempty"`
	// Enabled status of IPFIX export
	Enabled bool `json:"enabled"`
	// The maximum number of flow entries in each exporter flow cache. 
	MaxFlows int64 `json:"max_flows,omitempty"`
}