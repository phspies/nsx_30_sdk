package nsxt

// Used to provide the deployment specification for the service.
type ServiceDeployment struct {
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
	// This indicates the deployment perimeter, such as a VC cluster or a host.
	Perimeter string `json:"perimeter,omitempty"`
	// Name of the deployment spec to be used for deployment, which specifies the OVF provided by the partner and the form factor.
	DeploymentSpecName string `json:"deployment_spec_name"`
	// Mode of deployment. Currently, only stand alone deployment is supported. It is a single VM deployed through this deployment spec. In future, HA configurations will be supported here.
	DeploymentMode string `json:"deployment_mode,omitempty"`
	InstanceDeploymentTemplate *DeploymentTemplate `json:"instance_deployment_template"`
	ServiceDeploymentConfig *ServiceDeploymentConfig `json:"service_deployment_config"`
	// The Service to which the service deployment is associated.
	ServiceId string `json:"service_id,omitempty"`
	// Number of instances in case of clustered deployment.
	ClusteredDeploymentCount int64 `json:"clustered_deployment_count,omitempty"`
	// List of resource references where service instance be deployed. Ex. Tier 0 Logical Router in case of N-S ServiceInsertion. Service Attachment in case of E-W ServiceInsertion.
	DeployedTo []ResourceReference `json:"deployed_to,omitempty"`
	// Specifies whether the service VM should be deployed on each host such that it provides partner service locally on the host, or whether the service VMs can be deployed as a cluster. If deployment_type is CLUSTERED, then the clustered_deployment_count should be provided.
	DeploymentType string `json:"deployment_type,omitempty"`
}
