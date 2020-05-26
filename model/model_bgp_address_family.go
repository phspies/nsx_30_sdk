package nsxt

type BgpAddressFamily struct {
	// Count of out prefixes
	OutPrefixCount int64 `json:"out_prefix_count,omitempty"`
	// BGP address family type
	Type_ string `json:"type,omitempty"`
	// Count of in prefixes
	InPrefixCount int64 `json:"in_prefix_count,omitempty"`
}
