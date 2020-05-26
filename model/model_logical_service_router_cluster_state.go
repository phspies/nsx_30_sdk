package nsxt

// This holds the state of Service Router. If there are errors in realizing SR outside of MP, it gives details of the components and specific errors. 
type LogicalServiceRouterClusterState struct {
	// Request identifier of the API which modified the entity.
	PendingChangeList []string `json:"pending_change_list,omitempty"`
}
