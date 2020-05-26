package nsxt

// Registering a Service is the first step in the ServiceInsertion mechanism. A ServiceDefinition is used to create a service.
type ServiceDefinition struct {
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
	ServiceDeploymentSpec *ServiceDeploymentSpec `json:"service_deployment_spec,omitempty"`
	ServiceCapability *ServiceCapability `json:"service_capability,omitempty"`
	// The capabilities provided by the services. Needs to be one or more of the following | NG_FW - Next Generation Firewall | IDS_IPS - Intrusion detection System / Intrusion Prevention System | NET_MON - Network Monitoring | HCX - Hybrid Cloud Exchange | BYOD - Bring Your Own Device | EPP - Endpoint Protection.(Third party AntiVirus partners using NXGI should use this functionality for the service)
	Functionalities []string `json:"functionalities"`
	// The point at which the service is deployed/attached for redirecting the traffic to the the partner appliance. Attachment Point is required if Service caters to any functionality other than EPP.
	AttachmentPoint []string `json:"attachment_point,omitempty"`
	// ID of the service manager to which this service is attached with. This field is not set during creation of service. This field will be set explicitly when Service Manager is created successfully using this service. 
	ServiceManagerId string `json:"service_manager_id,omitempty"`
	// Id which is unique to a vendor or partner for which the service is created.
	VendorId string `json:"vendor_id"`
	// Failure policy for the service tells datapath, the action to take i.e to Allow or Block traffic during failure scenarios. For north-south ServiceInsertion, failure policy in the service instance takes precedence. For east-west ServiceInsertion, failure policy in the service chain takes precedence. BLOCK is not supported for Endpoint protection (EPP) functionality.
	OnFailurePolicy string `json:"on_failure_policy,omitempty"`
	// Transport Type of the service, which is the mechanism of redirecting the traffic to the the partner appliance. Transport type is required if Service caters to any functionality other than EPP.
	Transports []string `json:"transports,omitempty"`
	// This indicates the insertion point of the service i.e whether the service will be used to protect North-South or East-West traffic in the datacenter.
	Implementations []string `json:"implementations"`
}
