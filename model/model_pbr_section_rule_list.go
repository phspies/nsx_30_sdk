package nsxt

type PbrSectionRuleList struct {
	// Number of rules in this section.
	RuleCount int64 `json:"rule_count,omitempty"`
	// List of PBR rules in the section.
	Rules []PbrRule `json:"rules"`
}
