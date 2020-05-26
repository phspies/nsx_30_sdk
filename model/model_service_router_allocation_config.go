package nsxt

type ServiceRouterAllocationConfig struct {
	// For TIER 1 logical router, for manual placement of service router within the cluster, edge cluster member indices needs to be provided else same will be auto-allocated. You can provide maximum two indices for HA ACTIVE_STANDBY. 
	EdgeClusterMemberIndices []int64 `json:"edge_cluster_member_indices,omitempty"`
	AllocationPool *EdgeClusterMemberAllocationPool `json:"allocation_pool,omitempty"`
	// To reallocate TIER1 logical router on new or existing edge cluster 
	EdgeClusterId string `json:"edge_cluster_id"`
}
