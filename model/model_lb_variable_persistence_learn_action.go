package nsxt

// This action is performed in HTTP response rewrite phase. It is used to learn the value of variable from the HTTP response, and insert an entry into the persistence table if the entry doesn't exist. 
type LbVariablePersistenceLearnAction struct {
	// The property identifies the load balancer rule action type. 
	Type_ string `json:"type"`
	// The property is used to enable a hash operation for variable value when composing the persistence key. 
	VariableHashEnabled bool `json:"variable_hash_enabled,omitempty"`
	// The property is the name of variable to be learnt. It is used to identify which variable's value is learnt from HTTP response. The variable can be a system embedded variable such as \"_cookie_JSESSIONID\", a customized variable defined in LbVariableAssignmentAction or a captured variable in regular expression such as \"article\". 
	VariableName string `json:"variable_name"`
	// If the persistence profile UUID is not specified, a default persistence table is created per virtual server. Currently, only LbGenericPersistenceProfile is supported. 
	PersistenceProfileId string `json:"persistence_profile_id,omitempty"`
}