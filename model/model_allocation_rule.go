package nsxt

// Allocation rule on edge cluster which will be considered in auto placement of TIER1 logical routers, DHCP and MDProxy. 
type AllocationRule struct {
	Action *AllocationRuleAction `json:"action"`
}
