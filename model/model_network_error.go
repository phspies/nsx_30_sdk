package nsxt

// Network error related to container objects.
type NetworkError struct {
	// Detailed message of network related error.
	ErrorMessage string `json:"error_message,omitempty"`
	// Error code of network related error.
	ErrorCode string `json:"error_code,omitempty"`
	// Additional error information in json format.
	Spec string `json:"spec,omitempty"`
}
