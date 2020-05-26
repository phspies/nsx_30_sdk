package nsxt

// This action is used to rewrite header fields of HTTP response messages to specified new values at HTTP_RESPONSE_REWRITE phase. One action can be used to rewrite one header field. To rewrite multiple header fields, multiple actions must be defined. Captured variables and built-in variables can be used in the header_value field, header_name field does not support variables. 
type LbHttpResponseHeaderRewriteAction struct {
	// The property identifies the load balancer rule action type. 
	Type_ string `json:"type"`
	// Value of header field
	HeaderValue string `json:"header_value"`
	// Name of a header field of HTTP request message
	HeaderName string `json:"header_name"`
}
