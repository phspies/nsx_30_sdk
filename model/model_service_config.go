package nsxt

// Service configs are groupings of profiles (i.e switch profiles) and configurations applied to resources or collection of resources(NSGroups).
type ServiceConfig struct {
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
	// The list of entities that the configurations should be applied to. This can either be a NSGroup or any other entity like TransportNode, LogicalPorts etc. 
	AppliedTo []ResourceReference `json:"applied_to,omitempty"`
	// Every ServiceConfig has a priority based upon its precedence value. Lower the value of precedence, higher will be its priority. If user doesnt specify the precedence, it is generated automatically by system. The precedence is generated based upon the type of profile used in ServiceConfig. Precedence are auto-generated in decreasing order with difference of 100. Automatically generated precedence value will be 100 less than the current minimum value of precedence of ServiceConfig of a given profile type in system.There cannot be duplicate precedence for ServiceConfig of same profile type. 
	Precedence int64 `json:"precedence,omitempty"`
	// These are the NSX Profiles which will be added to service config, which will be applied to entities/groups provided to applied_to field of service config. 
	Profiles []NsxProfileReference `json:"profiles"`
}
