package nsxt

// Includes both scope and tag attribute of Tag. The scope and tag expressions are logically 'AND' with each other. eg- tag.scope = \"S1\" AND tag.tag = 'T1' 
type NsGroupTagExpression struct {
	ResourceType string `json:"resource_type"`
	// Target_type VirtualMachine supports all specified operators for tag expression while LogicalSwitch and LogicalPort supports only EQUALS operator. All operators perform a case insensitive match. 
	TagOp string `json:"tag_op,omitempty"`
	// The tag.scope attribute of the object
	Scope string `json:"scope,omitempty"`
	// Operator of the scope expression eg- tag.scope = \"S1\".
	ScopeOp string `json:"scope_op,omitempty"`
	// The tag.tag attribute of the object
	Tag string `json:"tag,omitempty"`
	// Type of the resource on which this expression is evaluated
	TargetType string `json:"target_type"`
}
