package nsxt

// Clustering parameters for the controller cluster
type ClusteringInfo struct {
	// Shared secret of the cluster.
	SharedSecret string `json:"shared_secret,omitempty"`
	// Property to indicate if the node must join an existing cluster.
	JoinToExistingCluster bool `json:"join_to_existing_cluster"`
}
