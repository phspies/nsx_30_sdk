package nsxt

// Contains a list of cluster node VM deployment requests and optionally a clustering configuration. 
type AddClusterNodeVmInfo struct {
	// Cluster node VM deployment requests to be deployed by the Manager. 
	DeploymentRequests []ClusterNodeVmDeploymentRequest `json:"deployment_requests"`
	ClusteringConfig *ClusteringConfig `json:"clustering_config,omitempty"`
}
