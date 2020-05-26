package nsxt

type LogicalPortOperationalStatus struct {
	// The id of the logical port
	LogicalPortId string `json:"logical_port_id,omitempty"`
	// The Operational status of the logical port
	Status string `json:"status"`
	// Timestamp when the data was last updated; unset if data source has never updated the data.
	LastUpdateTimestamp int64 `json:"last_update_timestamp,omitempty"`
}
