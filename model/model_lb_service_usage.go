package nsxt

// Describes the capacity and current usage of virtual servers, pools, pool members for the given load balancer service. 
type LbServiceUsage struct {
	// Pool capacity means maximum number of pools which could be configured in the given load balancer service. 
	PoolCapacity int64 `json:"pool_capacity,omitempty"`
	// The size of load balancer service
	ServiceSize string `json:"service_size,omitempty"`
	// The severity calculation is based on the largest usage percentage from virtual servers, pools, pool members and rules for one load balancer service. 
	Severity string `json:"severity,omitempty"`
	// Pool member capacity means maximum number of pool members which could be configured in the given load balancer service. 
	PoolMemberCapacity int64 `json:"pool_member_capacity,omitempty"`
	// The current number of virtual servers which have been configured in the given load balancer service. 
	CurrentVirtualServerCount int64 `json:"current_virtual_server_count,omitempty"`
	// The usage percentage is the largest usage percentage from virtual servers, pools and pool members for the load balancer service. If the property relax_scale_validation is set as true for LbService, it is possible that the value is larger than 100.0. For example, if SMALL LBS is deployed on MEDIUM edge node and configured with MEDIUM LBS virtual server scale number, LBS usage percentage is shown larger than 100.0. 
	UsagePercentage float64 `json:"usage_percentage,omitempty"`
	// UUID of load balancer service
	ServiceId string `json:"service_id,omitempty"`
	// The current number of pools which have been configured in the given load balancer service. 
	CurrentPoolCount int64 `json:"current_pool_count,omitempty"`
	// Virtual server capacity means maximum number of virtual servers which could be configured in the given load balancer service. 
	VirtualServerCapacity int64 `json:"virtual_server_capacity,omitempty"`
	// The current number of pool members which have been configured in the given load balancer service. 
	CurrentPoolMemberCount int64 `json:"current_pool_member_count,omitempty"`
}
