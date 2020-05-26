package nsxt

// Query statistics counters to an upstream server including successfully forwarded queries and failed queries. 
type PerUpstreamServerStatistics struct {
	// Queries failed to forward.
	QueriesFailed int64 `json:"queries_failed,omitempty"`
	// Upstream server ip
	UpstreamServer string `json:"upstream_server,omitempty"`
	// Queries forwarded successfully
	QueriesSucceeded int64 `json:"queries_succeeded,omitempty"`
}
