package nsxt

// List of ServiceInsertion Sections.
type ServiceInsertionSectionListResult struct {
	// List of the ServiceInsertion sections. The list has to be homogeneous.
	Results []ServiceInsertionSection `json:"results,omitempty"`
}
