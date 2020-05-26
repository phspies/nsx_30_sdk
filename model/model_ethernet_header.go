package nsxt

type EthernetHeader struct {
	// The destination MAC address of form: \"^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$\". For example: 00:00:00:00:00:00. 
	DstMac string `json:"dst_mac,omitempty"`
	// This field defaults to IPv4.
	EthType int64 `json:"eth_type,omitempty"`
	// The source MAC address of form: \"^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$\". For example: 00:00:00:00:00:00. 
	SrcMac string `json:"src_mac,omitempty"`
}
