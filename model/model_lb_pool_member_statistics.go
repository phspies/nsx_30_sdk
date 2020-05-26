package nsxt

type LbPoolMemberStatistics struct {
	Statistics *LbStatisticsCounter `json:"statistics"`
	// Pool member IP address
	IpAddress string `json:"ip_address"`
	// The port is configured in pool member. For virtual server port range case, pool member port must be null. 
	Port string `json:"port,omitempty"`
}
