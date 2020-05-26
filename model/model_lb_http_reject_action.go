package nsxt

// This action is used to reject HTTP request messages. The specified reply_status value is used as the status code for the corresponding HTTP response message which is sent back to client (Normally a browser) indicating the reason it was rejected. Reference official HTTP status code list for your specific HTTP version to set the reply_status properly. LbHttpRejectAction does not support variables. 
type LbHttpRejectAction struct {
	// The property identifies the load balancer rule action type. 
	Type_ string `json:"type"`
	// HTTP response status code
	ReplyStatus string `json:"reply_status"`
	// Response message
	ReplyMessage string `json:"reply_message,omitempty"`
}
