package nsxt

type PbrStats struct {
	Self *SelfResourceLink `json:"_self,omitempty"`
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink `json:"_links,omitempty"`
	// Schema for this resource
	Schema string `json:"_schema,omitempty"`
	// Aggregated number of packets processed by the rule.
	PacketCount int64 `json:"packet_count,omitempty"`
	// Aggregated number of bytes processed by the rule.
	ByteCount int64 `json:"byte_count,omitempty"`
	// Rule Identifier of the PBR rule. This is a globally unique number.
	RuleId string `json:"rule_id,omitempty"`
}
