package nsxt

// Details of infrastructure hosting the container cluster e.g. vSphere, AWS, VMC etc.. 
type ContainerInfrastructureInfo struct {
	// Type of the infrastructure.
	InfraType string `json:"infra_type"`
}
