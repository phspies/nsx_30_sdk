package nsxt

// This action is used to create a new variable and assign value to it. One action can be used to create one variable. To create multiple variables, multiple actions must be defined. The variables can be used by LbVariableCondition, etc. 
type LbVariableAssignmentAction struct {
	// The property identifies the load balancer rule action type. 
	Type_ string `json:"type"`
	// Value of variable
	VariableValue string `json:"variable_value"`
	// Name of the variable to be assigned
	VariableName string `json:"variable_name"`
}
