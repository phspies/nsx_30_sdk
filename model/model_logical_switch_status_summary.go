package nsxt

type LogicalSwitchStatusSummary struct {
	// Timestamp when the data was last updated; unset if data source has never updated the data.
	LastUpdateTimestamp int64 `json:"last_update_timestamp,omitempty"`
	// The total number of logical switches.
	TotalSwitches int64 `json:"total_switches"`
	// The filters used to find the logical switches- TransportZone id, LogicalSwitchProfile id or TransportType
	Filters []Filter `json:"filters,omitempty"`
	// The number of logical switches that are realized in all transport nodes.
	FullyRealizedSwitches int64 `json:"fully_realized_switches"`
}
