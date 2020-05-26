package nsxt

type LbServiceStatus struct {
	// status of load balancer pools
	Pools []LbPoolStatus `json:"pools,omitempty"`
	// Cpu usage in percentage
	CpuUsage int64 `json:"cpu_usage,omitempty"`
	// Ids of load balancer service related active transport nodes
	ActiveTransportNodes []string `json:"active_transport_nodes,omitempty"`
	// Memory usage in percentage
	MemoryUsage int64 `json:"memory_usage,omitempty"`
	// Load balancer service identifier
	ServiceId string `json:"service_id"`
	// Timestamp when the data was last updated
	LastUpdateTimestamp int64 `json:"last_update_timestamp,omitempty"`
	// Ids of load balancer service related standby transport nodes
	StandbyTransportNodes []string `json:"standby_transport_nodes,omitempty"`
	// Error message, if available
	ErrorMessage string `json:"error_message,omitempty"`
	// status of load balancer virtual servers
	VirtualServers []LbVirtualServerStatus `json:"virtual_servers,omitempty"`
	// UP means the load balancer service is working fine on both transport-nodes(if have); DOWN means the load balancer service is down on both transport-nodes (if have), hence the load balancer will not respond to any requests; ERROR means error happens on transport-node(s) or no status is reported from transport-node(s). The load balancer service may be working (or not working); NO_STANDBY means load balancer service is working in one of the transport node while not in the other transport-node (if have). Hence if the load balancer service in the working transport-node goes down, the load balancer service will go down; DETACHED means that the load balancer service has no attachment setting and is not instantiated in any transport nodes; DISABLED means that admin state of load balancer service is DISABLED; UNKNOWN means that no status reported from transport-nodes.The load balancer service may be working(or not working). 
	ServiceStatus string `json:"service_status,omitempty"`
}
