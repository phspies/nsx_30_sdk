package nsxt

// Details of the instructions displayed during restore process
type InstructionInfo struct {
	// A list of fields that are displayable to users in a table
	Fields []string `json:"fields,omitempty"`
	// UUID of the instruction
	Id string `json:"id,omitempty"`
	// A list of actions that are to be applied to resources
	Actions []string `json:"actions,omitempty"`
	// Instruction name
	Name string `json:"name,omitempty"`
}
