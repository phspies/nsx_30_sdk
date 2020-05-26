package nsxt

// Status of the Identity Firewall enabled Compute collection.
type IdfwComputeCollectionCondition struct {
	// IDFW enabled Compute collection status.
	Status string `json:"status"`
	// Status of the Compute collection.
	StatusDetail string `json:"status_detail,omitempty"`
}
