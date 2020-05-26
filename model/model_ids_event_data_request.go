package nsxt

// Filtering parameters to get only a subset of intrusion events.
type IdsEventDataRequest struct {
	// An array of filter conditions.
	Filters []FilterRequest `json:"filters,omitempty"`
}
