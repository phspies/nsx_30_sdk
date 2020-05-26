package nsxt

// A set of operations to be performed in a single batch
type BatchRequest struct {
	Requests []BatchRequestItem `json:"requests,omitempty"`
	// Flag to decide if we will continue processing subsequent requests in case of current error for atomic = false.
	ContinueOnError bool `json:"continue_on_error,omitempty"`
}
