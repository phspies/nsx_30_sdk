package nsxt

type LbPoolStatus struct {
	// UP means that all primary members are in UP status. PARTIALLY_UP means that some(not all) primary members are in UP status, the number of these active members is larger or equal to certain number(min_active_members) which is defined in LbPool. When there are no backup members which are in the UP status, the number(min_active_members) would be ignored. PRIMARY_DOWN means that less than certain(min_active_members) primary members are in UP status but backup members are in UP status, connections to this pool would be dispatched to backup members. DOWN means that all primary and backup members are DOWN. DETACHED means that the pool is not bound to any virtual server. UNKNOWN means that the pool is not associated to any enabled virtual servers, or no status reported from transport-nodes, the associated load balancer service may be working(or not working). 
	Status string `json:"status,omitempty"`
	// Timestamp when the data was last updated
	LastUpdateTimestamp int64 `json:"last_update_timestamp,omitempty"`
	// Load balancer pool identifier
	PoolId string `json:"pool_id"`
	// Status of load balancer pool members
	Members []LbPoolMemberStatus `json:"members,omitempty"`
}