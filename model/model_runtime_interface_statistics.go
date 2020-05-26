package nsxt

type RuntimeInterfaceStatistics struct {
	MacLearning *SiMacLearningCounters `json:"mac_learning,omitempty"`
	DroppedBySecurityPackets *SiPacketsDroppedBySecurity `json:"dropped_by_security_packets,omitempty"`
	// Timestamp when the data was last updated; unset if data source has never updated the data.
	LastUpdateTimestamp int64 `json:"last_update_timestamp,omitempty"`
	// Index of the interface
	InterfaceIndex int64 `json:"interface_index,omitempty"`
}
