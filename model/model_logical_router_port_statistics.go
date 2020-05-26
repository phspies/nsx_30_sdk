package nsxt

type LogicalRouterPortStatistics struct {
	// Per Node Statistics
	PerNodeStatistics []LogicalRouterPortStatisticsPerNode `json:"per_node_statistics,omitempty"`
	// The ID of the logical router port
	LogicalRouterPortId string `json:"logical_router_port_id"`
}
