package nsxt

// Runtime status information of the fabric node.
type NodeStatus struct {
	// Indicates the fabric node's MP&lt;-&gt;MPA channel connectivity status, UP, DOWN, UNKNOWN.
	MpaConnectivityStatus string `json:"mpa_connectivity_status,omitempty"`
	// Details, if any, about the current LCP&lt;-&gt;CCP channel connectivity status of the fabric node.
	LcpConnectivityStatusDetails []ControlConnStatus `json:"lcp_connectivity_status_details,omitempty"`
	// Details, if any, about the current MP&lt;-&gt;MPA channel connectivity status of the fabric node.
	MpaConnectivityStatusDetails string `json:"mpa_connectivity_status_details,omitempty"`
	// HostNode external id
	ExternalId string `json:"external_id,omitempty"`
	// Software version of the fabric node.
	SoftwareVersion string `json:"software_version,omitempty"`
	// Indicates the fabric node's status of maintenance mode, OFF, ENTERING, ON, EXITING.
	MaintenanceMode string `json:"maintenance_mode,omitempty"`
	// Is true if inventory sync is paused else false
	InventorySyncPaused bool `json:"inventory_sync_paused,omitempty"`
	SystemStatus *NodeStatusProperties `json:"system_status,omitempty"`
	// Inventory sync auto re-enable target time, in epoch milis
	InventorySyncReenableTime int64 `json:"inventory_sync_reenable_time,omitempty"`
	// Indicates the fabric node's LCP&lt;-&gt;CCP channel connectivity status, UP, DOWN, DEGRADED, UNKNOWN.
	LcpConnectivityStatus string `json:"lcp_connectivity_status,omitempty"`
	// Timestamp of the last heartbeat status change, in epoch milliseconds.
	LastHeartbeatTimestamp int64 `json:"last_heartbeat_timestamp,omitempty"`
	// Timestamp of the last successful update of Inventory, in epoch milliseconds.
	LastSyncTime int64 `json:"last_sync_time,omitempty"`
	// This enum specifies the current nsx install state for host node or current deployment and ready state for edge node. The ready status 'NODE_READY' indicates whether edge node is ready to become a transport node. The status 'EDGE_CONFIG_ERROR' indicates that edge hardware or underlying host is not supported. After all fabric level operations are done for an edge node, this value indicates transport node related configuration issues and state as relevant. 
	HostNodeDeploymentStatus string `json:"host_node_deployment_status,omitempty"`
}
