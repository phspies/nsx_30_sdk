package nsxt

type Ipv6Header struct {
	// The source ip address.
	SrcIp string `json:"src_ip,omitempty"`
	// The destination ip address.
	DstIp string `json:"dst_ip,omitempty"`
	// Identifies the type of header immediately following the IPv6 header.
	NextHeader int64 `json:"next_header,omitempty"`
	// Decremented by 1 by each node that forwards the packets. The packet is discarded if Hop Limit is decremented to zero.
	HopLimit int64 `json:"hop_limit,omitempty"`
}
