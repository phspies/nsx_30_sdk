package nsxt

type AggregatedDataCounter struct {
	TxBytes *DataCounter `json:"tx_bytes,omitempty"`
	RxPackets *DataCounter `json:"rx_packets,omitempty"`
	TxPackets *DataCounter `json:"tx_packets,omitempty"`
	RxBytes *DataCounter `json:"rx_bytes,omitempty"`
}
