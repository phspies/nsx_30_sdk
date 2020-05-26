package nsxt

// Community match operation
type CommunityMatchOperation struct {
	// Match operator for communities from provided community list id. MATCH_ANY will match any community MATCH_ALL will match all communities MATCH_EXACT will do exact match on community MATCH_NONE [operator not supported] will not match any community MATCH_REGEX will match normal communities by evaluating regular expression MATCH_LARGE_COMMUNITY_REGEX will match large communities by evaluating regular expression 
	MatchOperator string `json:"match_operator,omitempty"`
	// Regular expression to match BGP communities. If match_operator is MATCH_REGEX then this value must be specified. 
	RegularExpression string `json:"regular_expression,omitempty"`
	// ID of BGP community list. This value is not required when match_operator is MATCH_REGEX otherwise required. 
	CommunityListId string `json:"community_list_id,omitempty"`
}
