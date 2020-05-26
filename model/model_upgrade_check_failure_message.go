package nsxt

// Pre/post-upgrade check failure message
type UpgradeCheckFailureMessage struct {
	// Error/warning message
	Message string `json:"message,omitempty"`
	// Error code for the error/warning
	ErrorCode int64 `json:"error_code,omitempty"`
}
