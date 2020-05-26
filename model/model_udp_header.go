package nsxt

type UdpHeader struct {
	// Source port of udp header
	SrcPort int64 `json:"src_port,omitempty"`
	// Destination port of udp header
	DstPort int64 `json:"dst_port,omitempty"`
}
