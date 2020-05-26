package nsxt

type LbService struct {
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
	// Whether access log is enabled
	AccessLogEnabled bool `json:"access_log_enabled,omitempty"`
	Attachment *ResourceReference `json:"attachment,omitempty"`
	// Load balancer engine writes information about encountered issues of different severity levels to the error log. This setting is used to define the severity level of the error log. 
	ErrorLogLevel string `json:"error_log_level,omitempty"`
	// virtual servers can be associated to LbService(which is similar to physical/virtual load balancer), Lb virtual servers, pools and other entities could be defined independently, the virtual server identifier list here would be used to maintain the relationship of LbService and other Lb entities. 
	VirtualServerIds []string `json:"virtual_server_ids,omitempty"`
	// If relax_scale_validation is true, the scale validations for virtual servers/pools/pool members/rules are relaxed for load balancer service. When load balancer service is deployed on edge nodes, the scale of virtual servers/pools/pool members for the load balancer service should not exceed the scale number of the largest load balancer size which could be configured on a certain edge form factor. For example, the largest load balancer size supported on a MEDIUM edge node is MEDIUM. So one SMALL load balancer deployed on MEDIUM edge nodes can support the scale number of MEDIUM load balancer. It is not recommended to enable active monitors if relax_scale_validation is true due to performance consideration. If relax_scale_validation is false, scale numbers should be validated for load balancer service. 
	RelaxScaleValidation bool `json:"relax_scale_validation,omitempty"`
	// Whether the load balancer service is enabled
	Enabled bool `json:"enabled,omitempty"`
	// The size of load balancer service
	Size string `json:"size,omitempty"`
}
