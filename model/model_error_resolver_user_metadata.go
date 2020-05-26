package nsxt

// User supplied metadata needed for resolving errors
type ErrorResolverUserMetadata struct {
	// List of user supplied input data.
	UserInputList []ErrorResolverUserInputData `json:"user_input_list,omitempty"`
}
