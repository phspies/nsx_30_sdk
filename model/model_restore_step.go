package nsxt

// Restore step info
type RestoreStep struct {
	Status *PerStepRestoreStatus `json:"status,omitempty"`
	// Restore step number
	StepNumber int64 `json:"step_number,omitempty"`
	// Restore step description
	Description string `json:"description,omitempty"`
	// Restore step value
	Value string `json:"value,omitempty"`
}
