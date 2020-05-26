package nsxt

// Contains info used to configure the VM on deployment
type ClusterNodeVmDeploymentConfig struct {
	// Specifies the config for the platform through which to deploy the VM 
	PlacementType string `json:"placement_type"`
}
