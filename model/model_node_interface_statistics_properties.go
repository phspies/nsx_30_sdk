package nsxt

// Node network interface statistic properties
type NodeInterfaceStatisticsProperties struct {
	Self *SelfResourceLink `json:"_self,omitempty"`
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink `json:"_links,omitempty"`
	// Schema for this resource
	Schema string `json:"_schema,omitempty"`
	// Number of packets dropped
	TxDropped int64 `json:"tx_dropped,omitempty"`
	// Number of packets received
	RxPackets int64 `json:"rx_packets,omitempty"`
	// Number of carrier losses detected
	TxCarrier int64 `json:"tx_carrier,omitempty"`
	// Number of bytes received
	RxBytes int64 `json:"rx_bytes,omitempty"`
	// Number of transmit errors
	TxErrors int64 `json:"tx_errors,omitempty"`
	// Interface ID
	InterfaceId string `json:"interface_id,omitempty"`
	// Number of collisions detected
	TxColls int64 `json:"tx_colls,omitempty"`
	// Number of framing errors
	RxFrame int64 `json:"rx_frame,omitempty"`
	// Number of receive errors
	RxErrors int64 `json:"rx_errors,omitempty"`
	// Number of bytes transmitted
	TxBytes int64 `json:"tx_bytes,omitempty"`
	// Number of packets dropped
	RxDropped int64 `json:"rx_dropped,omitempty"`
	// Number of packets transmitted
	TxPackets int64 `json:"tx_packets,omitempty"`
	// Source of status data.
	Source string `json:"source,omitempty"`
}
