package nsxt

// Error information
type ErrorInfo struct {
	// Timestamp when the error occurred
	Timestamp int64 `json:"timestamp,omitempty"`
	// Error message
	ErrorMessage string `json:"error_message,omitempty"`
}
