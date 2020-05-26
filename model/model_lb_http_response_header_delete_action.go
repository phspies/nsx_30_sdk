package nsxt

// This action is used to delete header fields of HTTP response messages at HTTP_RESPONSE_REWRITE phase. One action can be used to delete allgi headers with same header name. To delete headers with different header names, multiple actions must be defined 
type LbHttpResponseHeaderDeleteAction struct {
	// The property identifies the load balancer rule action type. 
	Type_ string `json:"type"`
	// Name of a header field of HTTP response message
	HeaderName string `json:"header_name"`
}
