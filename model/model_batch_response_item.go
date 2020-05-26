package nsxt

// A single respose in a list of batched responses
type BatchResponseItem struct {
	// object returned by api
	Body *interface{} `json:"body,omitempty"`
	// The headers returned by the API call
	Headers *interface{} `json:"headers,omitempty"`
	// http status code
	Code int64 `json:"code"`
}
