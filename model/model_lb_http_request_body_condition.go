package nsxt

// This condition is used to match the message body of an HTTP request. Typically, only HTTP POST, PATCH, or PUT requests have request body. The match_type field defines how body_value field is used to match the body of HTTP requests. 
type LbHttpRequestBodyCondition struct {
	// A flag to indicate whether reverse the match result of this condition
	Inverse bool `json:"inverse,omitempty"`
	// Type of load balancer rule condition
	Type_ string `json:"type"`
	// HTTP request body
	BodyValue string `json:"body_value"`
	// Match type of HTTP body
	MatchType string `json:"match_type,omitempty"`
	// If true, case is significant when comparing HTTP body value. 
	CaseSensitive bool `json:"case_sensitive,omitempty"`
}
