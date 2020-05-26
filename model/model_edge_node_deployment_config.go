package nsxt

type EdgeNodeDeploymentConfig struct {
	NodeUserSettings *NodeUserSettings `json:"node_user_settings"`
	VmDeploymentConfig *DeploymentConfig `json:"vm_deployment_config"`
	// Supported edge form factor.
	FormFactor string `json:"form_factor,omitempty"`
}
