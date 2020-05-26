package nsxt

type MetadataProxyStatistics struct {
	// timestamp of the statistics
	Timestamp int64 `json:"timestamp"`
	// metadata proxy statistics per logical switch
	Statistics []MetadataProxyStatisticsPerLogicalSwitch `json:"statistics,omitempty"`
	// metadata proxy uuid
	MetadataProxyId string `json:"metadata_proxy_id"`
}
