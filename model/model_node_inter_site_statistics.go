package nsxt

type NodeInterSiteStatistics struct {
	// Timestamp when the remote tunnel port statistics was last updated. 
	LastUpdateTimestamp int64 `json:"last_update_timestamp,omitempty"`
	// Remote tunnel statistics per site.
	StatsPerSite []RemoteTunnelStatisticsPerSite `json:"stats_per_site,omitempty"`
	// Edge node id whose statistics is being reported.
	TransportNodeId string `json:"transport_node_id,omitempty"`
}
