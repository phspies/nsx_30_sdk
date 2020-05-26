package nsxt

// Error class for all the Service Insertion related errors.
type SiErrorClass struct {
	ErrorId int64 `json:"error_id,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`
}
