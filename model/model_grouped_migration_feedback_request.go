package nsxt

// Detailed feedback requests from the migration tool where user input is required.
type GroupedMigrationFeedbackRequest struct {
	// Indicates if a valid response already exist for all feedback requests in this group.
	Resolved bool `json:"resolved,omitempty"`
	// List of acceptable actions for this feedback request.
	AcceptedActions []string `json:"accepted_actions,omitempty"`
	// Identify a feedback request type across objects. This can be used to group together objects with similar feedback request and resolve them in one go.
	Hash string `json:"hash,omitempty"`
	// Functional area that this query falls into.
	Vertical string `json:"vertical,omitempty"`
	// The suggested value to resolve this feedback request.
	SuggestedValue string `json:"suggested_value,omitempty"`
	// Indicates if multiple values can be selected as response from the list of acceptable value.
	MultiValue bool `json:"multi_value,omitempty"`
	// Functional sub-area that this query falls into.
	SubVertical string `json:"sub_vertical,omitempty"`
	// Collection of feedback requests of a given type
	Objects []SummaryMigrationFeedbackRequest `json:"objects"`
	// List of acceptable values for this feedback request.
	AcceptedValues []string `json:"accepted_values,omitempty"`
	// Detailed feedback request with options.
	Message string `json:"message,omitempty"`
	// Data type of the items listed in acceptable values list.
	AcceptedValueType string `json:"accepted_value_type,omitempty"`
	// The suggested action to resolve this feedback request.
	SuggestedAction string `json:"suggested_action,omitempty"`
}