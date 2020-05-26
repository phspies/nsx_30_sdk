package nsxt

// Contains a list of NSX-Intelligence cluster node VM deployment requests and optionally a clustering configuration. 
type AddIntelligenceClusterNodeVmInfo struct {
	// Intelligence Cluster node VM deployment requests to be deployed by NSX. 
	DeploymentRequests []IntelligenceClusterNodeVmDeploymentRequest `json:"deployment_requests"`
}
