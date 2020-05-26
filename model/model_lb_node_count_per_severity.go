package nsxt

// The node count for specific load balancer usage severity. 
type LbNodeCountPerSeverity struct {
	// Node count for specific serverity. 
	NodeCount int64 `json:"node_count,omitempty"`
	// The severity calculation is based on credit usage percentage of load balancer for one node. 
	Severity string `json:"severity,omitempty"`
}