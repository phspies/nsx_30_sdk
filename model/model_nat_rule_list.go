package nsxt

type NatRuleList struct {
	// Add new NatRules to the list in Bulk creation. 
	Rules []NatRule `json:"rules"`
}
