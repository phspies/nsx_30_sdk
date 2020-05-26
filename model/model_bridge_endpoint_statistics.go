package nsxt

type BridgeEndpointStatistics struct {
	TxBytes *DataCounter `json:"tx_bytes,omitempty"`
	RxPackets *DataCounter `json:"rx_packets,omitempty"`
	TxPackets *DataCounter `json:"tx_packets,omitempty"`
	RxBytes *DataCounter `json:"rx_bytes,omitempty"`
	// Timestamp when the data was last updated; unset if data source has never updated the data.
	LastUpdateTimestamp int64 `json:"last_update_timestamp,omitempty"`
	// The id of the bridge endpoint
	EndpointId string `json:"endpoint_id,omitempty"`
}
