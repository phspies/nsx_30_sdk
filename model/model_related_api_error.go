package nsxt

// Detailed information about a related API error
type RelatedApiError struct {
	// The module name where the error occurred
	ModuleName string `json:"module_name,omitempty"`
	// A description of the error
	ErrorMessage string `json:"error_message,omitempty"`
	// A numeric error code
	ErrorCode int64 `json:"error_code,omitempty"`
	// Further details about the error
	Details string `json:"details,omitempty"`
	// Additional data about the error
	ErrorData *interface{} `json:"error_data,omitempty"`
}