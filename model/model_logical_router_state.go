package nsxt

// This holds the state of Logical Router. If there are errors in realizing LR outside of MP, it gives details of the components and specific errors.
type LogicalRouterState struct {
	// Request identifier of the API which modified the entity.
	PendingChangeList []string `json:"pending_change_list,omitempty"`
}
