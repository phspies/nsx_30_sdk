package nsxt

type FirewallSectionsSummary struct {
	Self *SelfResourceLink `json:"_self,omitempty"`
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink `json:"_links,omitempty"`
	// Schema for this resource
	Schema string `json:"_schema,omitempty"`
	// Total number of sections for the section type.
	SectionCount int64 `json:"section_count,omitempty"`
	// Total number of rules in the section.
	RuleCount int64 `json:"rule_count,omitempty"`
	// Type of rules which a section can contain.
	SectionType string `json:"section_type,omitempty"`
}
