package nsxt

// Complex expressions to represent NSGroup membership
type NsGroupComplexExpression struct {
	ResourceType string `json:"resource_type"`
	// Represents expressions which are to be logically 'AND'ed.The array cannot contain NSGroupComplexExpression.Only NSGroupTagExpression and NSGroupSimpleExpressions are accepted. 
	Expressions []NsGroupExpression `json:"expressions"`
}
