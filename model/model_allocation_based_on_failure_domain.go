package nsxt

// Auto place TIER1 logical routers, DHCP and MDProxy contexts on two edge nodes (active and standby) from different failure domains. 
type AllocationBasedOnFailureDomain struct {
	// Set action for each allocation rule on edge cluster which will help in auto placement. 
	ActionType string `json:"action_type"`
	// Enable placement algorithm to consider failure domain of edge transport nodes and place active and standby contexts in different failure domains. 
	Enabled bool `json:"enabled,omitempty"`
}
