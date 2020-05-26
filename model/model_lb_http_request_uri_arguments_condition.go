package nsxt

// This condition is used to match URI arguments aka query string of Http request messages, for example, in URI http://exaple.com?foo=1&bar=2, the \"foo=1&bar=2\" is the query string containing URI arguments. In an URI scheme, query string is indicated by the first question mark (\"?\") character and terminated by a number sign (\"#\") character or by the end of the URI. The uri_arguments field can be specified as a regular expression(Set match_type to REGEX). For example, \"foo=(?&lt;x&gt;\\d+)\". It matches HTTP requests whose URI arguments containing \"foo\", the value of foo contains only digits. And the value of foo is captured as $x which can be used in LbRuleAction fields which support variables. 
type LbHttpRequestUriArgumentsCondition struct {
	// A flag to indicate whether reverse the match result of this condition
	Inverse bool `json:"inverse,omitempty"`
	// Type of load balancer rule condition
	Type_ string `json:"type"`
	// URI arguments, aka query string of URI. 
	UriArguments string `json:"uri_arguments"`
	// Match type of URI arguments
	MatchType string `json:"match_type,omitempty"`
	// If true, case is significant when comparing URI arguments. 
	CaseSensitive bool `json:"case_sensitive,omitempty"`
}
