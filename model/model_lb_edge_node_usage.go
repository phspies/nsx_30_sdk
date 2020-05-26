package nsxt

// The capacity contains basic information and load balancer entity usages and capacity for the given edge node. 
type LbEdgeNodeUsage struct {
	// The property identifies the node UUID for load balancer node usage. 
	NodeId string `json:"node_id"`
	// The property identifies the load balancer node usage type. 
	Type_ string `json:"type"`
	// The current credit number reflects the current usage on the node. For example, configuring a medium load balancer on a node consumes 10 credits. If there are 2 medium instances configured on a node, the current credit number is 2 * 10 = 20. 
	CurrentCreditNumber int64 `json:"current_credit_number,omitempty"`
	// The form factor of the given edge node. 
	FormFactor string `json:"form_factor,omitempty"`
	// The number of virtual servers configured on the node. 
	CurrentVirtualServers int64 `json:"current_virtual_servers,omitempty"`
	// The number of small load balancer services configured on the node. 
	CurrentSmallLoadBalancerServices int64 `json:"current_small_load_balancer_services,omitempty"`
	// The number of pool members configured on the node. 
	CurrentPoolMembers int64 `json:"current_pool_members,omitempty"`
	// The severity calculation is based on current credit usage percentage of load balancer for one node. 
	Severity string `json:"severity,omitempty"`
	// The number of pools configured on the node. 
	CurrentPools int64 `json:"current_pools,omitempty"`
	// The remaining number of pool members which could be configured on the given edge node. 
	RemainingPoolMembers int64 `json:"remaining_pool_members,omitempty"`
	// The ID of edge cluster which contains the edge node. 
	EdgeClusterId string `json:"edge_cluster_id,omitempty"`
	// The remaining number of xlarge load balancer services which could be configured on the given edge node. 
	RemainingXlargeLoadBalancerServices int64 `json:"remaining_xlarge_load_balancer_services,omitempty"`
	// The remaining number of small load balancer services which could be configured on the given edge node. 
	RemainingSmallLoadBalancerServices int64 `json:"remaining_small_load_balancer_services,omitempty"`
	// The number of xlarge load balancer services configured on the node. 
	CurrentXlargeLoadBalancerServices int64 `json:"current_xlarge_load_balancer_services,omitempty"`
	// The usage percentage of the edge node for load balancer. The value is the larger value between load balancer credit usage percentage and pool member usage percentage for the edge node. 
	UsagePercentage float64 `json:"usage_percentage,omitempty"`
	// The number of large load balancer services configured on the node. 
	CurrentLargeLoadBalancerServices int64 `json:"current_large_load_balancer_services,omitempty"`
	// The remaining credit number is the remaining credits that can be used for load balancer service configuration. For example, an edge node with form factor LARGE_VIRTUAL_MACHINE has 40 credits, and a medium load balancer instance costs 10 credits. If there are currently 3 medium instances configured, the remaining credit number is 40 - (3 * 10) = 10. 
	RemainingCreditNumber int64 `json:"remaining_credit_number,omitempty"`
	// The remaining number of large load balancer services which could be configured on the given edge node. 
	RemainingLargeLoadBalancerServices int64 `json:"remaining_large_load_balancer_services,omitempty"`
	// The remaining number of medium load balancer services which could be configured on the given edge node. 
	RemainingMediumLoadBalancerServices int64 `json:"remaining_medium_load_balancer_services,omitempty"`
	// The number of medium load balancer services configured on the node. 
	CurrentMediumLoadBalancerServices int64 `json:"current_medium_load_balancer_services,omitempty"`
}
