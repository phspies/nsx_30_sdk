package nsxt

// The configuration entity to define a NAT rule. It defines how an ip packet is matched via source address or/and destination address or/and service(s), how the address (and/or) port is translated, and how the related firewall stage is involved or bypassed. 
type NatRule struct {
	// Indicates system owned resource
	SystemOwned bool `json:"_system_owned,omitempty"`
	// Defaults to ID if not set
	DisplayName string `json:"display_name,omitempty"`
	// Description of this resource
	Description string `json:"description,omitempty"`
	// Opaque identifiers meaningful to the API user
	Tags []Tag `json:"tags,omitempty"`
	// ID of the user who created this resource
	CreateUser string `json:"_create_user,omitempty"`
	// Protection status is one of the following: PROTECTED - the client who retrieved the entity is not allowed             to modify it. NOT_PROTECTED - the client who retrieved the entity is allowed                 to modify it REQUIRE_OVERRIDE - the client who retrieved the entity is a super                    user and can modify it, but only when providing                    the request header X-Allow-Overwrite=true. UNKNOWN - the _protection field could not be determined for this           entity. 
	Protection string `json:"_protection,omitempty"`
	// Timestamp of resource creation
	CreateTime int64 `json:"_create_time,omitempty"`
	// Timestamp of last modification
	LastModifiedTime int64 `json:"_last_modified_time,omitempty"`
	// ID of the user who last modified this resource
	LastModifiedUser string `json:"_last_modified_user,omitempty"`
	// Unique identifier of this resource
	Id string `json:"id,omitempty"`
	// The type of this resource.
	ResourceType string `json:"resource_type,omitempty"`
	// IP Address | CIDR | (null implies Any) 
	MatchDestinationNetwork string `json:"match_destination_network,omitempty"`
	// The translated address for the matched IP packet. For a SNAT, it can be a single ip address, an ip range, or a CIDR block. For a DNAT and a REFLEXIVE, it can be a single ip address or a CIDR block. Translated network is not supported for NO_SNAT or NO_DNAT. 
	TranslatedNetwork string `json:"translated_network,omitempty"`
	// Ascending, valid range [0-2147483647]. If multiple rules have the same priority, evaluation sequence is undefined. 
	RulePriority int64 `json:"rule_priority,omitempty"`
	MatchService *NsServiceElement `json:"match_service,omitempty"`
	// Holds the list of LogicalRouterPort Ids that a NAT rule can be applied to. The LogicalRouterPort used must belong to the same LogicalRouter for which the NAT Rule is created. As of now a NAT rule can only have a single LogicalRouterPort as applied_tos. When applied_tos is not set, the NAT rule is applied to all LogicalRouterPorts beloging to the LogicalRouter.
	AppliedTos []ResourceReference `json:"applied_tos,omitempty"`
	// Indicator to enable/disable the rule. 
	Enabled bool `json:"enabled,omitempty"`
	// Internal NAT rule uuid for debug used in Controller and backend.
	InternalRuleId string `json:"internal_rule_id,omitempty"`
	// Enable/disable the logging of rule. 
	Logging bool `json:"logging,omitempty"`
	// The translated port(s) for the mtached IP packet. It can be a single port or a port range. Please note, port translating is supported only for DNAT. 
	TranslatedPorts string `json:"translated_ports,omitempty"`
	// Valid actions: SNAT, DNAT, NO_SNAT, NO_DNAT, REFLEXIVE, NAT64. All rules in a logical router are either stateless or stateful. Mix is not supported. SNAT and DNAT are stateful, can NOT be supported when the logical router is running at active-active HA mode; REFLEXIVE is stateless. NO_SNAT and NO_DNAT have no translated_fields, only match fields are supported. 
	Action string `json:"action"`
	// Indicate how firewall is applied to a traffic packet. Firewall can be bypassed, or be applied to external/internal address of NAT rule.  The firewall_match will take priority over nat_pass. If the firewall_match is not provided, the nat_pass will be picked up. 
	FirewallMatch string `json:"firewall_match,omitempty"`
	// Default is true. If the nat_pass is set to true, the following firewall stage will be skipped. Please note, if action is NO_SNAT or NO_DNAT, then nat_pass must be set to true or omitted.  Nat_pass was deprecated with an alternative firewall_match. Please stop using nat_pass to specify whether firewall stage is skipped. if you want to skip, please set firewall_match to BYPASS. If you do not want to skip, please set the firewall_match to MATCH_EXTERNAL_ADDRESS or MATCH_INTERNAL_ADDRESS.  Please note, the firewall_match will take priority over the nat_pass. If both are provided, the nat_pass is ignored. If firewall_match is not provided while the nat_pass is specified, the nat_pass will still be picked up. In this case, if nat_pass is set to false, firewall rule will be applied on internall address of a packet, i.e. MATCH_INTERNAL_ADDRESS. 
	NatPass bool `json:"nat_pass,omitempty"`
	// The logical router id which the nat rule runs on.
	LogicalRouterId string `json:"logical_router_id,omitempty"`
	// IP Address | CIDR | (null implies Any) 
	MatchSourceNetwork string `json:"match_source_network,omitempty"`
}