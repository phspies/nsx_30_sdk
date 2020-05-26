package nsxt

type AdvanceClusterRestoreInput struct {
	// Unique id of an instruction (as returned by the GET /restore/status call) for which input is to be provided 
	Id string `json:"id,omitempty"`
	// List of resources for which the instruction is applicable.
	Resources []SelectableResourceReference `json:"resources"`
}
