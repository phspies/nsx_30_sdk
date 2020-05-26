package nsxt

type SiMacLearningCounters struct {
	// Number of MACs learned
	MacsLearned int64 `json:"macs_learned,omitempty"`
	// The number of packets with unknown source MAC address that are dropped without learning the source MAC address. Applicable only when the MAC limit is reached and MAC Limit policy is MAC_LEARNING_LIMIT_POLICY_DROP.
	MacNotLearnedPacketsDropped int64 `json:"mac_not_learned_packets_dropped,omitempty"`
	// The number of packets with unknown source MAC address that are dispatched without learning the source MAC address. Applicable only when the MAC limit is reached and MAC Limit policy is MAC_LEARNING_LIMIT_POLICY_ALLOW.
	MacNotLearnedPacketsAllowed int64 `json:"mac_not_learned_packets_allowed,omitempty"`
}
