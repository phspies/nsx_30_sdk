package nsxt

// Routing policy details.
type RoutingPolicy struct {
	// Array of next hop to prefix lists mapping.
	NextHopPrefixListsMappings []NextHopPrefixListsMapping `json:"next_hop_prefix_lists_mappings"`
	// Routing policy type.
	RoutingPolicyType string `json:"routing_policy_type,omitempty"`
}
