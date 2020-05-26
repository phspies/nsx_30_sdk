package nsxt

// Node usage for load balancer contains basic information and LB entity usages and capacity for the given node. 
type LbNodeUsage struct {
	// The property identifies the node UUID for load balancer node usage. 
	NodeId string `json:"node_id"`
	// The property identifies the load balancer node usage type. 
	Type_ string `json:"type"`
}
