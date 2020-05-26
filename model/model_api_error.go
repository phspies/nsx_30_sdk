package nsxt

// Detailed information about an API Error
type ApiError struct {
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
	// Other errors related to this error
	RelatedErrors []RelatedApiError `json:"related_errors,omitempty"`
}
