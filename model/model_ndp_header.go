package nsxt

// Neighbor discovery protocol header
type NdpHeader struct {
	// The IP address of the destination of the solicitation. It MUST NOT be a multicast address.
	DstIp string `json:"dst_ip,omitempty"`
	// This field specifies the type of the Neighbor discover message being sent. NEIGHBOR_SOLICITATION - Neighbor Solicitation message to discover the link-layer address of an on-link IPv6 node or to confirm a previously determined link-layer address. NEIGHBOR_ADVERTISEMENT - Neighbor Advertisement message in response to a Neighbor Solicitation message.
	MsgType string `json:"msg_type,omitempty"`
}
