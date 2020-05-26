package nsxt

// The current statistics counters of the DNS forwarder including cache usages and query numbers per forwarders. 
type DnsForwarderStatistics struct {
	// The total number of forwarded dns queries
	QueriesForwarded int64 `json:"queries_forwarded,omitempty"`
	// The statistics of conditional forwarders
	ConditionalForwarderStatistics []PerForwarderStatistics `json:"conditional_forwarder_statistics,omitempty"`
	DefaultForwarderStatistics *PerForwarderStatistics `json:"default_forwarder_statistics,omitempty"`
	// The totocal number of queries answered from local cache
	QueriesAnsweredLocally int64 `json:"queries_answered_locally,omitempty"`
	// The statistics of used cache
	UsedCacheStatistics []PerNodeUsedCacheStatistics `json:"used_cache_statistics,omitempty"`
	// The configured cache size, in kb
	ConfiguredCacheSize int64 `json:"configured_cache_size,omitempty"`
	// Time stamp of the current statistics, in ms
	Timestamp int64 `json:"timestamp,omitempty"`
	// Error message, if available
	ErrorMessage string `json:"error_message,omitempty"`
	// The total number of received dns queries
	TotalQueries int64 `json:"total_queries,omitempty"`
}
