package nsxt

// Restore step status
type PerStepRestoreStatus struct {
	// A description of the restore status
	Description string `json:"description,omitempty"`
	// Per step restore status value
	Value string `json:"value,omitempty"`
}
