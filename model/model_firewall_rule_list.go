package nsxt

type FirewallRuleList struct {
	// List of firewall rules in the section. Only homogenous rules are supported.
	Rules []FirewallRule `json:"rules"`
}
