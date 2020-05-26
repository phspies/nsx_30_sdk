package nsxt

type IcmpEchoRequestHeader struct {
	// ICMP id
	Id int64 `json:"id,omitempty"`
	// ICMP sequence number
	Sequence int64 `json:"sequence,omitempty"`
}
