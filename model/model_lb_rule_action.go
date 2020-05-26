package nsxt

// Load balancer rule actions are used to manipulate application traffic. Currently load balancer rules can be used at three load balancer processing phases. Each phase has its own supported type of actions. Supported actions in HTTP_REQUST_REWRITE phase are:   LbHttpRequestUriRewriteAction   LbHttpRequestHeaderRewriteAction   LbHttpRequestHeaderDeleteAction   LbVariableAssignmentAction Supported actions in HTTP_FORWARDING phase are:   LbHttpRejectAction   LbHttpRedirectAction   LbSelectPoolAction   LbVariablePersistenceOnAction   LbConnectionDropAction Supported action in HTTP_RESPONSE_REWRITE phase is:   LbHttpResponseHeaderRewriteAction   LbHttpResponseHeaderDeleteAction   LbVariablePersistenceLearnAction Supported action in HTTP_ACCESS phase is:   LbJwtAuthAction   LbConnectionDropAction   LbVariableAssignmentAction Supported action in TRANSPORT phase is:   LbSslModeSelectionAction   LbSelectPoolAction  If the match type of an LbRuleCondition field is specified as REGEX and named capturing groups are used in the specified regular expression. The groups can be used as variables in LbRuleAction fields. For example, define a rule with LbHttpRequestUriCondition as match condition and LbHttpRequestUriRewriteAction as action. Set match_type field of LbHttpRequestUriCondition to REGEX, and set uri field to   \"/news/(?&lt;year&gt;\\d+)/(?&lt;month&gt;\\d+)/(?&lt;article&gt;.*)\". Set uri field of LbHttpRequestUriRewriteAction to:   \"/news/$year-$month/$article\" In uri field of LbHttpRequestUriCondition, the (?&lt;year&gt;\\d+), (?&lt;month&gt;\\d+) and (?&lt;article&gt;.*) are named capturing groups, they define variables named $year, $month and $article respectively. The defined variables are used in LbHttpRequestUriRewriteAction. For a matched HTTP request with URI \"/news/2017/06/xyz.html\", the substring \"2017\" is captured in variable $year, \"06\" is captured in variable $month, and \"xyz.html\" is captured in variable $article. The LbHttpRequestUriRewriteAction will rewrite the URI to:   \"/news/2017-06/xyz.html\" A set of built-in variables can be used in LbRuleAction as well. The name of built-in variables start with underscore, the name of user defined variables is not allowed to start with underscore. Following are some of the built-in variables:   $_scheme: Reference the scheme part of matched HTTP messages, could be \"http\" or \"https\".   $_host: Host of matched HTTP messages, for example \"www.example.com\".   $_server_port: Port part of URI, it is also the port of the server which accepted a request. Default port is 80 for http and 443 for https.   $_uri: The URI path, for example \"/products/sample.html\".   $_request_uri: Full original request URI with arguments, for example,     \"/products/sample.html?a=b&c=d\".   $_args: URI arguments, for instance \"a=b&c=d\"   $_is_args: \"?\" if a request has URI arguments, or an empty string otherwise. For the full list of built-in variables, please reference the NSX-T Administrator's Guide. 
type LbRuleAction struct {
	// The property identifies the load balancer rule action type. 
	Type_ string `json:"type"`
}
