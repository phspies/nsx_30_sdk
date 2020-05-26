package nsxt

type BridgeClusterStatus struct {
	// Timestamp when the data was last updated; unset if data source has never updated the data.
	LastUpdateTimestamp int64 `json:"last_update_timestamp,omitempty"`
	// The health status of the cluster
	Health string `json:"health,omitempty"`
	// The id of the cluster
	ClusterId string `json:"cluster_id,omitempty"`
}
