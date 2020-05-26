package nsxt

// This condition is used to match variable's name and value at all phases. The variables could be captured from REGEX or assigned by LbVariableAssignmentAction or system embedded variable. Varialbe_name and variable_value should be matched at the same time. 
type LbVariableCondition struct {
	// A flag to indicate whether reverse the match result of this condition
	Inverse bool `json:"inverse,omitempty"`
	// Type of load balancer rule condition
	Type_ string `json:"type"`
	// Value of variable to be matched
	VariableValue string `json:"variable_value"`
	// If true, case is significant when comparing variable value. 
	CaseSensitive bool `json:"case_sensitive,omitempty"`
	// Match type of variable value
	MatchType string `json:"match_type,omitempty"`
	// Name of the variable to be matched
	VariableName string `json:"variable_name"`
}
