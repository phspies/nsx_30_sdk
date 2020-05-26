package nsxt

// List of ServiceInsertion Rules.
type ServiceInsertionRuleList struct {
	// List of ServiceInsertion rules in the section. Only homogeneous rules are supported.
	Rules []ServiceInsertionRule `json:"rules"`
}
