package nsxt

type SiAggregatedDataCounterEx struct {
	TxBytes *SiDataCounter `json:"tx_bytes,omitempty"`
	RxPackets *SiDataCounter `json:"rx_packets,omitempty"`
	TxPackets *SiDataCounter `json:"tx_packets,omitempty"`
	RxBytes *SiDataCounter `json:"rx_bytes,omitempty"`
	MacLearning *SiMacLearningCounters `json:"mac_learning,omitempty"`
	DroppedBySecurityPackets *SiPacketsDroppedBySecurity `json:"dropped_by_security_packets,omitempty"`
}
