package nsxt

// A Service Runtime is the runtime entity associated with ever Service-VM deployed.
type InstanceRuntime struct {
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
	// Service-VM/SVM id of deployed virtual-machine.
	ServiceVmId string `json:"service_vm_id,omitempty"`
	// Service-Instance Runtime deployment status of the Service-VM. It shows the latest status during the process of deployment, redeploy, upgrade, and un-deployment of VM.
	DeploymentStatus string `json:"deployment_status,omitempty"`
	VmNicInfo *VmNicInfo `json:"vm_nic_info,omitempty"`
	// The maintenance mode indicates whether the corresponding service VM is in maintenance mode. The service VM will not be used to service new requests if it is in maintenance mode. 
	MaintenanceMode string `json:"maintenance_mode,omitempty"`
	// Service-Instance Runtime status of the deployed Service-VM.
	RuntimeStatus string `json:"runtime_status,omitempty"`
	// Error message for the Service Instance Runtime if any.
	ErrorMessage string `json:"error_message,omitempty"`
	// Id of an instantiation of a registered service.
	ServiceInstanceId string `json:"service_instance_id,omitempty"`
	// Service-Instance runtime health status set by partner to indicate whether the service is running properly or not. 
	RuntimeHealthStatusByPartner string `json:"runtime_health_status_by_partner,omitempty"`
	// Reason provided by partner for the service being unhealthy. This could be due to various reasons such as connectivity lost as an example. 
	UnhealthyReason string `json:"unhealthy_reason,omitempty"`
}
