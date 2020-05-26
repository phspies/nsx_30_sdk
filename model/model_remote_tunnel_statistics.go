package nsxt

type RemoteTunnelStatistics struct {
	// Ip address of remote tunnel destination.
	TunnelDestinationAddress string `json:"tunnel_destination_address,omitempty"`
	// Ip address of remote tunnel source.
	TunnelSourceAddress string `json:"tunnel_source_address,omitempty"`
	Rx *InterSitePortCounters `json:"rx,omitempty"`
	Tx *InterSitePortCounters `json:"tx,omitempty"`
}
