package nsxt

type ArpHeader struct {
	// This field specifies the nature of the Arp message being sent.
	OpCode string `json:"op_code"`
	// This field specifies the IP address of the sender. If omitted, the src_ip is set to 0.0.0.0.
	SrcIp string `json:"src_ip,omitempty"`
	// The destination IP address
	DstIp string `json:"dst_ip"`
}
