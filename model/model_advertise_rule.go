package nsxt

type AdvertiseRule struct {
	// ALLOW action enables the advertisment and DENY action disables the advertisement of a filtered routes to the connected TIER0 router.
	Action string `json:"action,omitempty"`
	RuleFilter *AdvertisementRuleFilter `json:"rule_filter,omitempty"`
	// Display name
	DisplayName string `json:"display_name,omitempty"`
	// network(CIDR) to be routed
	Networks []string `json:"networks"`
	// Description
	Description string `json:"description,omitempty"`
}
