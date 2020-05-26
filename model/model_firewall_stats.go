package nsxt

type FirewallStats struct {
	Self *SelfResourceLink `json:"_self,omitempty"`
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink `json:"_links,omitempty"`
	// Schema for this resource
	Schema string `json:"_schema,omitempty"`
	// Aggregated number of sessions processed by the all firewall rules. This is aggregated statistic which are computed with lower frequency compared to individual generic rule statistics. It may have a computation delay up to 15 minutes in response to this API.
	TotalSessionCount int64 `json:"total_session_count,omitempty"`
	// Aggregated number of packets processed by the rule.
	PacketCount int64 `json:"packet_count,omitempty"`
	// This is calculated by sessions count divided by age of the rule.
	PopularityIndex int64 `json:"popularity_index,omitempty"`
	// Maximum value of sessions count of all firewall rules of the type. This is aggregated statistic which are computed with lower frequency compared to generic rule statistics. It may have a computation delay up to 15 minutes in response to this API.
	MaxSessionCount int64 `json:"max_session_count,omitempty"`
	// Aggregated number of bytes processed by the rule.
	ByteCount int64 `json:"byte_count,omitempty"`
	// Maximum value of popularity index of all firewall rules of the type. This is aggregated statistic which are computed with lower frequency compared to individual generic rule statistics. It may have a computation delay up to 15 minutes in response to this API.
	MaxPopularityIndex int64 `json:"max_popularity_index,omitempty"`
	// Aggregated number of hits received by the rule.
	HitCount int64 `json:"hit_count,omitempty"`
	// Aggregated number of sessions processed by the rule.
	SessionCount int64 `json:"session_count,omitempty"`
	// Rule Identifier of the Firewall rule. This is a globally unique number.
	RuleId string `json:"rule_id,omitempty"`
}
