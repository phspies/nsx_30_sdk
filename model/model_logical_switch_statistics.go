package nsxt

type LogicalSwitchStatistics struct {
	MacLearning *MacLearningCounters `json:"mac_learning,omitempty"`
	DroppedBySecurityPackets *PacketsDroppedBySecurity `json:"dropped_by_security_packets,omitempty"`
	// Timestamp when the data was last updated; unset if data source has never updated the data.
	LastUpdateTimestamp int64 `json:"last_update_timestamp,omitempty"`
	// The id of the logical Switch
	LogicalSwitchId string `json:"logical_switch_id,omitempty"`
}
