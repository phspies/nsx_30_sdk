package nsxt

// This condition is used to match URIs(Uniform Resource Identifier) of HTTP request messages. The URI field can be specified as a regular expression. If an HTTP request message is requesting an URI which matches specified regular expression, it matches the condition. The syntax of whole URI looks like this: scheme:[//[user[:password]@]host[:port]][/path][?query][#fragment] This condition matches only the path part of entire URI. When match_type field is specified as REGEX, the uri field is used as a regular expression to match URI path of HTTP requests. For example, to match any URI that has \"/image/\" or \"/images/\", uri field can be specified as: \"/image[s]?/\". Named capturing groups can be used in the uri field to capture substrings of matched URIs and store them in variables for use in LbRuleAction. For example, specify uri field as: \"/news/(?&lt;year&gt;\\d+)/(?&lt;month&gt;\\d+)/(?&lt;article&gt;.*)\" If the URI path is /articles/news/2017/06/xyz.html, then substring \"2017\" is captured in variable year, \"06\" is captured in variable month, and \"xyz.html\" is captured in variable article. These variables can then be used in an LbRuleAction field which supports variables, such as uri field of LbHttpRequestUriRewriteAction. For example, set the uri field of LbHttpRequestUriRewriteAction as: \"/articles/news/$year-$month-$article\" Then the URI path /articles/news/2017/06/xyz.html is rewritten to: \"/articles/news/2017-06-xyz.html\" 
type LbHttpRequestUriCondition struct {
	// A flag to indicate whether reverse the match result of this condition
	Inverse bool `json:"inverse,omitempty"`
	// Type of load balancer rule condition
	Type_ string `json:"type"`
	// Match type of URI
	MatchType string `json:"match_type,omitempty"`
	// A string used to identify resource
	Uri string `json:"uri"`
	// If true, case is significant when comparing URI. 
	CaseSensitive bool `json:"case_sensitive,omitempty"`
}
