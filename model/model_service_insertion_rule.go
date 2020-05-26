package nsxt

// ServiceInsertion Rule.
type ServiceInsertionRule struct {
	// Flag to indicate whether rule is default.
	IsDefault bool `json:"is_default,omitempty"`
	// Rule direction in case of stateless distributed service rules. This will only considered if section level parameter is set to stateless. Default to IN_OUT if not specified.
	Direction string `json:"direction,omitempty"`
	// User level field which will be printed in CLI and packet logs.
	RuleTag string `json:"rule_tag,omitempty"`
	// Type of IP packet that should be matched while enforcing the rule.
	IpProtocol string `json:"ip_protocol,omitempty"`
	// User notes specific to the rule.
	Notes string `json:"notes,omitempty"`
	// List of object where rule will be enforced. The section level field overrides this one. Null will be treated as any.
	AppliedTos []ResourceReference `json:"applied_tos,omitempty"`
	// Flag to enable packet logging. Default is disabled.
	Logged bool `json:"logged,omitempty"`
	// Flag to disable rule. Disabled will only be persisted but never provisioned/realized.
	Disabled bool `json:"disabled,omitempty"`
	// List of sources. Null will be treated as any.
	Sources []ResourceReference `json:"sources,omitempty"`
	// Action enforced on the packets which matches the distributed service rule. Currently DS Layer supports below actions. ALLOW           - Forward any packet when a rule with this action gets a match (Used by Firewall). DROP            - Drop any packet when a rule with this action gets a match. Packets won't go further(Used by Firewall). REJECT          - Terminate TCP connection by sending TCP reset for a packet when a rule with this action gets a match (Used by Firewall). REDIRECT        - Redirect any packet to a partner appliance when a rule with this action gets a match (Used by Service Insertion). DO_NOT_REDIRECT - Do not redirect any packet to a partner appliance when a rule with this action gets a match (Used by Service Insertion). DETECT          - Detect IDS Signatures.
	Action string `json:"action"`
	// Priority of the rule.
	Priority int64 `json:"priority,omitempty"`
	// Negation of the source.
	SourcesExcluded bool `json:"sources_excluded,omitempty"`
	// Negation of the destination.
	DestinationsExcluded bool `json:"destinations_excluded,omitempty"`
	// List of the destinations. Null will be treated as any.
	Destinations []ResourceReference `json:"destinations,omitempty"`
	// List of the services. Null will be treated as any.
	Services []ServiceInsertionService `json:"services,omitempty"`
	// A rule can be redirected to ServiceInstance, InstanceEndpoint for North/South Traffic. A rule can be redirected to ServiceChain for East/West Traffic. For REDIRECT action, redirect_tos is mandatory. For DO_NOT_REDIRECT action, redirect_tos is optional.
	RedirectTos []ResourceReference `json:"redirect_tos,omitempty"`
	// ID of the section to which this rule belongs.
	SectionId string `json:"section_id,omitempty"`
}
