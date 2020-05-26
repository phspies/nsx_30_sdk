package nsxt

type ServiceInstanceStatus struct {
	InstanceDeploymentStatus *ServiceDeploymentStatus `json:"instance_deployment_status,omitempty"`
	ConfigurationIssue *SvmConfigureIssue `json:"configuration_issue,omitempty"`
	// Id of an instantiation of a registered service.
	ServiceInstanceId string `json:"service_instance_id,omitempty"`
	InstanceHealthStatus *ServiceInstanceHealthStatus `json:"instance_health_status,omitempty"`
}
