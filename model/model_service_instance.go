package nsxt

// The deployment of a registered service. Service instance is instantiation of service. It is the most used type of instance. It is a default instance to be used when NSX handles lifecycle of appliance. Deployment and appliance related all the information is necessary.
type ServiceInstance struct {
	// Failure policy of the service instance - if it has to be different from the service. By default the service instance inherits the FailurePolicy of the service it belongs to.
	OnFailurePolicy string `json:"on_failure_policy,omitempty"`
	// Transport to be used by this service instance for deploying the Service-VM. This field is to be set Not Applicable(NA) if the service only caters to functionality EPP(Endpoint Protection).
	TransportType string `json:"transport_type"`
	// ServiceInstance is used when NSX handles the lifecyle of   appliance. Deployment and appliance related all the information is necessary. ByodServiceInstance is a custom instance to be used when NSX is not handling   the lifecycles of appliance/s. User will manage their own appliance (BYOD)   to connect with NSX. VirtualServiceInstance is a a custom instance to be used when NSX is not   handling the lifecycle of an appliance and when the user is not bringing   their own appliance. 
	ResourceType string `json:"resource_type"`
	// The Service to which the service instance is associated.
	ServiceId string `json:"service_id,omitempty"`
	// Name of the deployment spec to be used by this service instance.
	DeploymentSpecName string `json:"deployment_spec_name"`
	InstanceDeploymentTemplate *DeploymentTemplate `json:"instance_deployment_template"`
	// Implementation to be used by this service instance for deploying the Service-VM.
	ImplementationType string `json:"implementation_type"`
	// Attachment point to be used by this service instance for deploying the Service-VM.
	AttachmentPoint string `json:"attachment_point"`
	InstanceDeploymentConfig *InstanceDeploymentConfig `json:"instance_deployment_config,omitempty"`
	// Deployment mode specifies where the partner appliance will be deployed in HA or non-HA i.e standalone mode.
	DeploymentMode string `json:"deployment_mode"`
	// List of resource references where service instance be deployed. Ex. Tier 0 Logical Router in case of N-S ServiceInsertion.
	DeployedTo []ResourceReference `json:"deployed_to"`
	// Id of the Service Deployment using which the instances were deployed. Its available only for instances that were deployed using service deployment API.
	ServiceDeploymentId string `json:"service_deployment_id,omitempty"`
}
