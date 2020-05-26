package nsxt

// This action is used to select a pool for matched HTTP request messages. The pool is specified by UUID. The matched HTTP request messages are forwarded to the specified pool. 
type LbSelectPoolAction struct {
	// The property identifies the load balancer rule action type. 
	Type_ string `json:"type"`
	// Display name of load balancer pool
	PoolName string `json:"pool_name,omitempty"`
	// UUID of load balancer pool
	PoolId string `json:"pool_id"`
}
