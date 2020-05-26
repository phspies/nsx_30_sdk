package nsxt

// This action is used to redirect HTTP request messages to a new URL. The reply_status value specified in this action is used as the status code of HTTP response message which is sent back to client (Normally a browser). The HTTP status code for redirection is 3xx, for example, 301, 302, 303, 307, etc. The redirect_url is the new URL that the HTTP request message is redirected to. Normally browser will send another HTTP request to the new URL after receiving a redirection response message. Captured variables and built-in variables can be used in redirect_url field. For example, to redirect all HTTP requests to HTTPS requests for a virtual server. We create an LbRule without any conditions, add an LbHttpRedirectAction to the rule. Set the redirect_url field of the LbHttpRedirectAction to:   https://$_host$_request_uri And set redirect_status to \"302\", which means found. This rule will redirect all HTTP requests to HTTPS server port on the same host. 
type LbHttpRedirectAction struct {
	// The property identifies the load balancer rule action type. 
	Type_ string `json:"type"`
	// HTTP response status code
	RedirectStatus string `json:"redirect_status"`
	// The URL that the HTTP request is redirected to
	RedirectUrl string `json:"redirect_url"`
}
