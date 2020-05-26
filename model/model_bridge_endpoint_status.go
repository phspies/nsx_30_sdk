package nsxt

type BridgeEndpointStatus struct {
	// The Ids of the transport nodes which actively serve the endpoint.
	ActiveNodes []string `json:"active_nodes,omitempty"`
	// Timestamp when the data was last updated; unset if data source has never updated the data.
	LastUpdateTimestamp int64 `json:"last_update_timestamp,omitempty"`
	// The id of the bridge endpoint
	EndpointId string `json:"endpoint_id,omitempty"`
}
