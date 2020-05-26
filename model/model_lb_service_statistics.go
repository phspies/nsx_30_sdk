package nsxt

type LbServiceStatistics struct {
	// Statistics of load balancer pools
	Pools []LbPoolStatistics `json:"pools,omitempty"`
	// load balancer service identifier
	ServiceId string `json:"service_id"`
	// Statistics of load balancer virtual servers
	VirtualServers []LbVirtualServerStatistics `json:"virtual_servers,omitempty"`
	// Timestamp when the data was last updated
	LastUpdateTimestamp int64 `json:"last_update_timestamp,omitempty"`
	Statistics *LbServiceStatisticsCounter `json:"statistics,omitempty"`
}
