package nsxt

// The array of the failed DNS queries with entry count and timestamp on active and standby transport node. 
type DnsFailedQueries struct {
	// Timestamp of the request, in YYYY-MM-DD HH:MM:SS.zzz format. 
	Timestamp string `json:"timestamp,omitempty"`
	// The array of failed DNS queries on active and standby transport node. If there is no standby node, the failed queries on standby node will not be present. 
	PerNodeFailedQueries []PerNodeDnsFailedQueries `json:"per_node_failed_queries,omitempty"`
}
