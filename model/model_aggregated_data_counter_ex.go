package nsxt

type AggregatedDataCounterEx struct {
	TxBytes *DataCounter `json:"tx_bytes,omitempty"`
	RxPackets *DataCounter `json:"rx_packets,omitempty"`
	TxPackets *DataCounter `json:"tx_packets,omitempty"`
	RxBytes *DataCounter `json:"rx_bytes,omitempty"`
	MacLearning *MacLearningCounters `json:"mac_learning,omitempty"`
	DroppedBySecurityPackets *PacketsDroppedBySecurity `json:"dropped_by_security_packets,omitempty"`
}
