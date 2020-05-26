package nsxt

// This condition is used to match the HTTP protocol version of the HTTP request messages. 
type LbHttpRequestVersionCondition struct {
	// A flag to indicate whether reverse the match result of this condition
	Inverse bool `json:"inverse,omitempty"`
	// Type of load balancer rule condition
	Type_ string `json:"type"`
	// HTTP version
	Version string `json:"version"`
}
