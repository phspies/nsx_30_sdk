package nsxt

// Detailed feedback requests from the migration tool where user input is required.
type MigrationFeedbackResponse struct {
	// Action selected in response to the feedback request.
	Action string `json:"action"`
	// User input provided in the form of a list of values in response to the feedback request.
	Values []string `json:"values,omitempty"`
	// Identifier of the feedback request.
	Id string `json:"id"`
	// User input provided in response to the feedback request.
	Value string `json:"value,omitempty"`
}