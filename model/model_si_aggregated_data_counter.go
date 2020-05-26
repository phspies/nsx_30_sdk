package nsxt

type SiAggregatedDataCounter struct {
	TxBytes *SiDataCounter `json:"tx_bytes,omitempty"`
	RxPackets *SiDataCounter `json:"rx_packets,omitempty"`
	TxPackets *SiDataCounter `json:"tx_packets,omitempty"`
	RxBytes *SiDataCounter `json:"rx_bytes,omitempty"`
}
