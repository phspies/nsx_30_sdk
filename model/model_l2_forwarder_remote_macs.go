package nsxt

type L2ForwarderRemoteMacs struct {
	// Timestamp when the l2 forwarder remote mac addresses was last updated. 
	LastUpdateTimestamp int64 `json:"last_update_timestamp,omitempty"`
	// Logical switch id on which the L2 forwarder is created. 
	LogicalSwitchId string `json:"logical_switch_id,omitempty"`
	// L2 forwarder remote mac addresses per site for logical switch. 
	RemoteMacsPerSite []L2ForwarderRemoteMacsPerSite `json:"remote_macs_per_site,omitempty"`
}
