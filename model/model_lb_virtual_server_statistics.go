package nsxt

type LbVirtualServerStatistics struct {
	// Timestamp when the data was last updated
	LastUpdateTimestamp int64 `json:"last_update_timestamp,omitempty"`
	Statistics *LbStatisticsCounter `json:"statistics"`
	// load balancer virtual server identifier
	VirtualServerId string `json:"virtual_server_id"`
}
