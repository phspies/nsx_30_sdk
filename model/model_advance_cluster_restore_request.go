package nsxt

type AdvanceClusterRestoreRequest struct {
	// List of instructions and their associated data
	Data []AdvanceClusterRestoreInput `json:"data"`
}
