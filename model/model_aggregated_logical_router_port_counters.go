package nsxt

type AggregatedLogicalRouterPortCounters struct {
	// Timestamp when the data was last updated; unset if data source has never updated the data.
	LastUpdateTimestamp int64 `json:"last_update_timestamp,omitempty"`
	Rx *LogicalRouterPortCounters `json:"rx,omitempty"`
	Tx *LogicalRouterPortCounters `json:"tx,omitempty"`
}
