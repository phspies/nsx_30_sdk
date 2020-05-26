package nsxt

// Categorization of feedback requests from the migration tool where user input is required.
type MigrationFeedbackCategory struct {
	// Functional area that this query falls into.
	Category string `json:"category,omitempty"`
	// Total number of feedback requests for this functional area.
	Count int32 `json:"count,omitempty"`
	// Total number of resolved feedback requests for this functional area.
	Resolved int32 `json:"resolved,omitempty"`
	// List of acceptable values for this feedback request.
	AcceptedValues []string `json:"accepted_values,omitempty"`
}
