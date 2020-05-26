package nsxt

type L2ForwarderStatistics struct {
	// Timestamp when the l2 forwarder statistics was last updated. 
	LastUpdateTimestamp int64 `json:"last_update_timestamp,omitempty"`
	// Logical switch id on which the L2 forwarder is created. 
	LogicalSwitchId string `json:"logical_switch_id,omitempty"`
	Rx *InterSitePortCounters `json:"rx,omitempty"`
	Tx *InterSitePortCounters `json:"tx,omitempty"`
}
