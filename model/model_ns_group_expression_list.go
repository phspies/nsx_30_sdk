package nsxt

// List of NSGroupExpressions
type NsGroupExpressionList struct {
	// List of NSGroupExpressions to be passed to add and remove APIs 
	Members []NsGroupExpression `json:"members"`
}
