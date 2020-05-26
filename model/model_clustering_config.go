package nsxt

// Configuration for automatically joining a cluster node to the cluster after it is deployed. ClusteringConfig is required if any of the deployment nodes has CONTROLLER role. 
type ClusteringConfig struct {
	// Specifies the type of clustering config to be used. 
	ClusteringType string `json:"clustering_type"`
}
