package nsxt

// The load balancer node usage summary for all nodes. Only EdgeNode is supported. The summary calculation is based on all edge nodes configured in edge clusters. 
type LbNodeUsageSummary struct {
	// The property identifies array of node count for each severity (RED, ORANGE and GREEN). 
	NodeCounts []LbNodeCountPerSeverity `json:"node_counts,omitempty"`
	// The current credit number reflects the overall credit usage for all nodes. 
	CurrentCreditNumber int64 `json:"current_credit_number,omitempty"`
	// The property contains lb node usages for each node. 
	NodeUsages []LbNodeUsage `json:"node_usages,omitempty"`
	// The severity calculation is based on current credit usage percentage of load balancer for all nodes. 
	Severity string `json:"severity,omitempty"`
	// The overall remaining number of pool members which could be configured on all nodes. 
	RemainingPoolMembers int64 `json:"remaining_pool_members,omitempty"`
	// The overall number of pool members configured on all nodes. 
	CurrentPoolMembers int64 `json:"current_pool_members,omitempty"`
	// The overall usage percentage of all nodes for load balancer. The value is the larger value between overall pool member usage percentage and overall load balancer credit usage percentage. 
	UsagePercentage float64 `json:"usage_percentage,omitempty"`
	// The remaining credit number is the overall remaining credits that can be used for load balancer service configuration for all nodes. 
	RemainingCreditNumber int64 `json:"remaining_credit_number,omitempty"`
}
