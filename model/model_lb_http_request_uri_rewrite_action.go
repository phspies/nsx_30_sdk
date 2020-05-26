package nsxt

// This action is used to rewrite URIs in matched HTTP request messages. Specify the uri and uri_arguments fields in this condition to rewrite the matched HTTP request message's URI and URI arguments to the new values. Full URI scheme of HTTP messages have following syntax: scheme:[//[user[:password]@]host[:port]][/path][?query][#fragment] The uri field of this action is used to rewrite the /path part in above scheme. And the uri_arguments field is used to rewrite the query part. Captured variables and built-in variables can be used in the uri and uri_arguments fields. Check the example in LbRuleAction to see how to use variables in this action. 
type LbHttpRequestUriRewriteAction struct {
	// The property identifies the load balancer rule action type. 
	Type_ string `json:"type"`
	// Query string of URI, typically contains key value pairs, for example: foo1=bar1&foo2=bar2 
	UriArguments string `json:"uri_arguments,omitempty"`
	// URI of HTTP request
	Uri string `json:"uri"`
}
