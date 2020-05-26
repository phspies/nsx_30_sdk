package nsxt

// Allocation details of cluster and its members. Contains information of the edge nodes present in cluster, active and standby services of each node, utilization details of configured sub-pools. These allocation details can be monitored by customers to trigger migration of certain service contexts to different edge nodes, to balance the utilization of edge node resources. 
type EdgeClusterAllocationStatus struct {
	// Allocation details of edge nodes present in the cluster.
	Members []EdgeMemberAllocationStatus `json:"members,omitempty"`
	// Display name of the edge cluster
	DisplayName string `json:"display_name,omitempty"`
	// System allotted UUID of edge cluster.
	Id string `json:"id,omitempty"`
	// Represents the number of edge nodes in the cluster.
	MemberCount int32 `json:"member_count,omitempty"`
}
