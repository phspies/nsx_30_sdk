package nsxt

// This action is used to delete header fields of HTTP request messages at HTTP_REQUEST_REWRITE phase. One action can be used to delete all headers with same header name. To delete headers with different header names, multiple actions must be defined. 
type LbHttpRequestHeaderDeleteAction struct {
	// The property identifies the load balancer rule action type. 
	Type_ string `json:"type"`
	// Name of a header field of HTTP request message
	HeaderName string `json:"header_name"`
}
