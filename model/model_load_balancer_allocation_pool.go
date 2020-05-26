package nsxt

type LoadBalancerAllocationPool struct {
	// Types of logical router allocation pool based on services
	AllocationPoolType string `json:"allocation_pool_type"`
	// To address varied customer performance and scalability requirements, different sizes for load balancer service are supported: SMALL, MEDIUM, LARGE and XLARGE, each with its own set of resource and performance. Specify size of load balancer service which you will bind to TIER1 router. 
	AllocationSize string `json:"allocation_size"`
}
