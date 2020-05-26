package nsxt

// Contains the deployment information for a cluster node VM soon to be deployed or already deployed by the Manager 
type ClusterNodeVmDeploymentRequest struct {
	DeploymentConfig *ClusterNodeVmDeploymentConfig `json:"deployment_config"`
	// ID of the VM maintained internally and used to recognize it. Note: This is automatically generated and cannot be modified. 
	VmId string `json:"vm_id,omitempty"`
	UserSettings *NodeUserSettings `json:"user_settings,omitempty"`
	// List of cluster node role (or roles) which the VM should take on. They specify what type (or types) of cluster node which the new VM should act as. Currently both CONTROLLER and MANAGER must be provided, since this permutation is the only one supported now. 
	Roles []string `json:"roles"`
	// Specifies the desired \"size\" of the VM 
	FormFactor string `json:"form_factor,omitempty"`
}
