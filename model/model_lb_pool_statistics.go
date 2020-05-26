package nsxt

type LbPoolStatistics struct {
	// Timestamp when the data was last updated
	LastUpdateTimestamp int64 `json:"last_update_timestamp,omitempty"`
	Statistics *LbStatisticsCounter `json:"statistics"`
	// Load balancer pool identifier
	PoolId string `json:"pool_id"`
	// Statistics of load balancer pool members
	Members []LbPoolMemberStatistics `json:"members,omitempty"`
}
