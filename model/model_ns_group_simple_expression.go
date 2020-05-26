package nsxt

// Simple expressions to represent NSGroup membership
type NsGroupSimpleExpression struct {
	ResourceType string `json:"resource_type"`
	TargetResource *ResourceReference `json:"target_resource,omitempty"`
	// Field of the resource on which this expression is evaluated
	TargetProperty string `json:"target_property"`
	// Type of the resource on which this expression is evaluated
	TargetType string `json:"target_type"`
	// Value that satisfies this expression
	Value string `json:"value"`
	// All operators perform a case insensitive match. 
	Op string `json:"op"`
}
