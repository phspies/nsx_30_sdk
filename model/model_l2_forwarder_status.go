package nsxt

type L2ForwarderStatus struct {
	// Timestamp when the service router status was last updated. 
	LastUpdateTimestamp int64 `json:"last_update_timestamp,omitempty"`
	// Logical switch id on which the L2 forwarder is created. 
	LogicalSwitchId string `json:"logical_switch_id,omitempty"`
	// L2 forwarder status per node.
	StatusPerNode []L2ForwarderStatusPerNode `json:"status_per_node,omitempty"`
}
