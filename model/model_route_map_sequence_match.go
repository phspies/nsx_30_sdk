package nsxt

// Match sequence in route map which is used for matching routes. IP prefix lists and match community expression are mutually exclusive fields, one of them must be provided. 
type RouteMapSequenceMatch struct {
	MatchCommunityExpression *CommunityMatchExpression `json:"match_community_expression,omitempty"`
	// IPPrefixList Identifiers for RouteMap Sequence Match Criteria
	IpPrefixLists []string `json:"ip_prefix_lists,omitempty"`
}
