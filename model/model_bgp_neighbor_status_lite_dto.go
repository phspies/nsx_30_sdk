package nsxt

type BgpNeighborStatusLiteDto struct {
	// Current state of the BGP session.
	ConnectionState string `json:"connection_state,omitempty"`
	RemoteSite *ResourceReference `json:"remote_site,omitempty"`
	// Source Ip address.
	SourceAddress string `json:"source_address,omitempty"`
	// Ip address of BGP neighbor.
	NeighborAddress string `json:"neighbor_address,omitempty"`
}
