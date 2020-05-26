package nsxt

// The reponse to a batch operation
type BatchResponse struct {
	// Indicates if any of the APIs failed
	HasErrors bool `json:"has_errors,omitempty"`
	// Optional flag indicating that all items were rolled back even if succeeded initially
	RolledBack bool `json:"rolled_back,omitempty"`
	// Bulk list results
	Results []BatchResponseItem `json:"results"`
}
