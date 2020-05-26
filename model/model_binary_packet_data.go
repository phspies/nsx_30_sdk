package nsxt

type BinaryPacketData struct {
	// A flag, when set true, indicates that the traceflow packet is of L3 routing.
	Routed bool `json:"routed,omitempty"`
	// transport type of the traceflow packet
	TransportType string `json:"transport_type,omitempty"`
	// Packet configuration
	ResourceType string `json:"resource_type"`
	// If the requested frame_size is too small (given the payload and traceflow metadata requirement of 16 bytes), the traceflow request will fail with an appropriate message.  The frame will be zero padded to the requested size.
	FrameSize int64 `json:"frame_size,omitempty"`
	// Up to 1000 bytes of payload may be supplied (with a base64-encoded length of 1336 bytes.) Additional bytes of traceflow metadata will be appended to the payload. The payload must contain all headers (Ethernet, IP, etc). Note that VLAN is not supported in the logical space. Hence, payload must not contain 802.1Q headers.
	Payload string `json:"payload,omitempty"`
}