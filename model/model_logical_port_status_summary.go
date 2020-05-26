package nsxt

type LogicalPortStatusSummary struct {
	// The total number of logical ports.
	TotalPorts int64 `json:"total_ports"`
	// Timestamp when the data was last updated; unset if data source has never updated the data.
	LastUpdateTimestamp int64 `json:"last_update_timestamp,omitempty"`
	// The number of logical ports whose Operational status is UP
	UpPorts int64 `json:"up_ports"`
	// The filters used to find the logical ports- TransportZone id, LogicalSwitch id or LogicalSwitchProfile id
	Filters []Filter `json:"filters,omitempty"`
}
