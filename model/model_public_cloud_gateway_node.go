package nsxt

type PublicCloudGatewayNode struct {
	NodeSettings *EdgeNodeSettings `json:"node_settings,omitempty"`
	DeploymentConfig *EdgeNodeDeploymentConfig `json:"deployment_config,omitempty"`
	// List of logical router ids to which this edge node is allocated.
	AllocationList []string `json:"allocation_list,omitempty"`
	// Supported edge deployment type.
	DeploymentType string `json:"deployment_type,omitempty"`
}
