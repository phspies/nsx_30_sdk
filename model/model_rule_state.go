package nsxt

type RuleState struct {
	// This attribute represents revision number of rule's desired state.
	RevisionDesired int64 `json:"revision_desired,omitempty"`
	// Pending changes to be realized.
	PendingChangeList []PendingChange `json:"pending_change_list,omitempty"`
}
