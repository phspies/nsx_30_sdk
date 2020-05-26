package nsxt

// This condition is used to match HTTP response messages from backend servers by HTTP header fields. HTTP header fields are components of the header section of HTTP request and response messages. They define the operating parameters of an HTTP transaction. For example, Cookie, Authorization, User-Agent, etc. One condition can be used to match one header field, to match multiple header fields, multiple conditions must be specified. The match_type field defines how header_value field is used to match HTTP responses. The header_name field does not support match types. 
type LbHttpResponseHeaderCondition struct {
	// A flag to indicate whether reverse the match result of this condition
	Inverse bool `json:"inverse,omitempty"`
	// Type of load balancer rule condition
	Type_ string `json:"type"`
	// Value of HTTP header field
	HeaderValue string `json:"header_value"`
	// If true, case is significant when comparing HTTP header value. 
	CaseSensitive bool `json:"case_sensitive,omitempty"`
	// Match type of HTTP header value
	MatchType string `json:"match_type,omitempty"`
	// Name of HTTP header field
	HeaderName string `json:"header_name"`
}