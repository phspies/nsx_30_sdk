package nsxt

type ManagementNodeAggregateInfo struct {
	// Array of Node interface statistic properties
	NodeInterfaceProperties []NodeInterfaceProperties `json:"node_interface_properties,omitempty"`
	NodeStatus *ClusterNodeStatus `json:"node_status,omitempty"`
	// Array of Node network interface statistic properties
	NodeInterfaceStatistics []NodeInterfaceStatisticsProperties `json:"node_interface_statistics,omitempty"`
	// Defaults to ID if not set
	DisplayName string `json:"display_name,omitempty"`
	// Time series of the node's system properties
	NodeStatusProperties []NodeStatusProperties `json:"node_status_properties,omitempty"`
	// Unique identifier of this resource
	Id string `json:"id,omitempty"`
	TransportNodesConnected int64 `json:"transport_nodes_connected,omitempty"`
	RoleConfig *ManagementClusterRoleConfig `json:"role_config,omitempty"`
}
