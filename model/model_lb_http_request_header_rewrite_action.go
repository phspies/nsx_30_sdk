package nsxt

// This action is used to rewrite header fields of matched HTTP request messages to specified new values. One action can be used to rewrite one header field. To rewrite multiple header fields, multiple actions must be defined. Captured variables and built-in variables can be used in the header_value field, header_name field does not support variables. 
type LbHttpRequestHeaderRewriteAction struct {
	// The property identifies the load balancer rule action type. 
	Type_ string `json:"type"`
	// Value of HTTP request header
	HeaderValue string `json:"header_value"`
	// Name of HTTP request header
	HeaderName string `json:"header_name"`
}
