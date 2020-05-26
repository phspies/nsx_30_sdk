package nsxt

type NatStatisticsPerLogicalRouter struct {
	// Timestamp when the data was last updated; unset if data source has never updated the data.
	LastUpdateTimestamp int64 `json:"last_update_timestamp,omitempty"`
	// Detailed per node statistics
	PerTransportNodeStatistics []NatStatisticsPerTransportNode `json:"per_transport_node_statistics,omitempty"`
	StatisticsAcrossAllNodes *NatCounters `json:"statistics_across_all_nodes,omitempty"`
	// Id for the logical router
	LogicalRouterId string `json:"logical_router_id,omitempty"`
}
