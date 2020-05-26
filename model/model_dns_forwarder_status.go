package nsxt

// The current runtime status of the DNS forwarder including the hosting transport nodes and forwarder service status. 
type DnsForwarderStatus struct {
	// UP means the DNS forwarder is working correctly on the active transport node and the stand-by transport node (if present). Failover will occur if either node goes down. DOWN means the DNS forwarder is down on both active transport node and standby node (if present). The DNS forwarder does not function in this situation. Error means there is some error on one or both transport node, or no status was reported from one or both transport nodes. The dns forwarder may be working (or not working). NO_BACKUP means dns forwarder is working in only one transport node, either because it is down on the standby node, or no standby is configured. An forwarder outage will occur if the active node goes down. 
	Status string `json:"status,omitempty"`
	// Time stamp of the current status, in ms
	Timestamp int64 `json:"timestamp,omitempty"`
	// Uuid of stand_by transport node. null if non-HA mode
	StandbyNode string `json:"standby_node,omitempty"`
	// Extra message, if available
	ExtraMessage string `json:"extra_message,omitempty"`
	// Uuid of active transport node
	ActiveNode string `json:"active_node,omitempty"`
}