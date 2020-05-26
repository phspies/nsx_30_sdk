package nsxt

// Contains the deployment information for a NSX-Intelligence node VM. 
type IntelligenceClusterNodeVmDeploymentRequest struct {
	DeploymentConfig *IntelligenceClusterNodeVmDeploymentConfig `json:"deployment_config"`
	// ID of the VM maintained internally. Note: This is automatically generated and cannot be modified. 
	VmId string `json:"vm_id,omitempty"`
	UserSettings *NodeUserSettings `json:"user_settings,omitempty"`
	// Specifies the desired \"size\" of the VM 
	FormFactor string `json:"form_factor,omitempty"`
}
