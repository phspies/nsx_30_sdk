package nsxt

type LogicalRouterStatus struct {
	// Timestamp when the data was last updated; unset if data source has never updated the data.
	LastUpdateTimestamp int64 `json:"last_update_timestamp,omitempty"`
	// The id of the logical router
	LogicalRouterId string `json:"logical_router_id"`
	// Per Node Status
	PerNodeStatus []LogicalRouterStatusPerNode `json:"per_node_status,omitempty"`
	// Egress mode for the logical router at given mode 
	LocaleOperationMode string `json:"locale_operation_mode,omitempty"`
}
