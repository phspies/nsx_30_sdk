package nsxt

type TcpHeader struct {
	// TCP flags (9bits)
	TcpFlags int64 `json:"tcp_flags,omitempty"`
	// Source port of tcp header
	SrcPort int64 `json:"src_port,omitempty"`
	// Destination port of tcp header
	DstPort int64 `json:"dst_port,omitempty"`
}
