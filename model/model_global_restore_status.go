package nsxt

// Overall restore process status
type GlobalRestoreStatus struct {
	// A description of the restore status
	Description string `json:"description,omitempty"`
	// Global rolled-up restore status value
	Value string `json:"value,omitempty"`
}
