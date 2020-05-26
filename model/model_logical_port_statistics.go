package nsxt

type LogicalPortStatistics struct {
	MacLearning *MacLearningCounters `json:"mac_learning,omitempty"`
	DroppedBySecurityPackets *PacketsDroppedBySecurity `json:"dropped_by_security_packets,omitempty"`
	// The id of the logical port
	LogicalPortId string `json:"logical_port_id,omitempty"`
	// Timestamp when the data was last updated; unset if data source has never updated the data.
	LastUpdateTimestamp int64 `json:"last_update_timestamp,omitempty"`
}
