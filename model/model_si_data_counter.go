package nsxt

type SiDataCounter struct {
	// The total packets or bytes
	Total int64 `json:"total"`
	// The multicast and broadcast packets or bytes
	MulticastBroadcast int64 `json:"multicast_broadcast,omitempty"`
	// The dropped packets or bytes
	Dropped int64 `json:"dropped,omitempty"`
}
