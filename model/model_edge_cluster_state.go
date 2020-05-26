package nsxt

// This holds the state of Edge Cluster. If there are errors in realizing EC outside of MP, it gives details of the components and specific errors. 
type EdgeClusterState struct {
	// Request identifier of the API which modified the entity.
	PendingChangeList []string `json:"pending_change_list,omitempty"`
}
