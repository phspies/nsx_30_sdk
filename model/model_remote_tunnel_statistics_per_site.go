package nsxt

type RemoteTunnelStatisticsPerSite struct {
	// Statistics per remote tunnel.
	StatsPerTunnel []RemoteTunnelStatistics `json:"stats_per_tunnel,omitempty"`
	RemoteSite *ResourceReference `json:"remote_site,omitempty"`
	Rx *InterSitePortCounters `json:"rx,omitempty"`
	Tx *InterSitePortCounters `json:"tx,omitempty"`
}
