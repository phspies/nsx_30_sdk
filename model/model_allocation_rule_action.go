package nsxt

// Define action for each allocation rule which added on edge cluster. 
type AllocationRuleAction struct {
	// Set action for each allocation rule on edge cluster which will help in auto placement. 
	ActionType string `json:"action_type"`
}
