package nsxt

type EdgeMemberAllocationStatus struct {
	// List of services allocated on the edge node.
	AllocatedServices []AllocatedService `json:"allocated_services,omitempty"`
	// Display name of edge cluster member. Defaults to ID if not set. 
	NodeDisplayName string `json:"node_display_name,omitempty"`
	// System generated index for transport node backed by edge node. 
	MemberIndex int32 `json:"member_index,omitempty"`
	// Allocation details of pools defined on the edge node.
	AllocationPools []AllocationPool `json:"allocation_pools,omitempty"`
	// System allotted UUID of edge node.
	NodeId string `json:"node_id,omitempty"`
}
