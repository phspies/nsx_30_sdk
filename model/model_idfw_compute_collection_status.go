package nsxt

// compute collection ID and status connected to VC.
type IdfwComputeCollectionStatus struct {
	// IDFW enabled compute collection status.
	ComputeCollectionStatus []IdfwComputeCollectionCondition `json:"compute_collection_status,omitempty"`
	// IDFW compute collection ID connected to VC.
	ComputeCollectionId string `json:"compute_collection_id"`
}
