package nsxt

type SiPacketTypeAndCounter struct {
	// The number of packets.
	Counter int64 `json:"counter"`
	// The type of the packets
	PacketType string `json:"packet_type"`
}
