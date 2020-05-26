package nsxt

// Community match expression
type CommunityMatchExpression struct {
	// Operator for evaluating community match expressions. AND logical AND operator 
	Operator string `json:"operator,omitempty"`
	// Array of community match operations
	Expression []CommunityMatchOperation `json:"expression"`
}
