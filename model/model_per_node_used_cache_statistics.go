package nsxt

// Query statistics counters of used cache from node 
type PerNodeUsedCacheStatistics struct {
	// The total number of cached entries
	CachedEntries int64 `json:"cached_entries,omitempty"`
	// Uuid of active/standby transport node
	NodeId string `json:"node_id,omitempty"`
	// The memory size used in cache, in kb
	UsedCacheSize int64 `json:"used_cache_size,omitempty"`
}
