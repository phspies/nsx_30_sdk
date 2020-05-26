package nsxt

type TransportNodeStatus struct {
	// Roll-up status of pNIC, management connection, control connection, tunnel status, agent status
	Status string `json:"status,omitempty"`
	ThreatStatus *ThreatStatus `json:"threat_status,omitempty"`
	AgentStatus *AgentStatusCount `json:"agent_status,omitempty"`
	// Transport node uuid
	NodeUuid string `json:"node_uuid,omitempty"`
	TunnelStatus *TunnelStatusCount `json:"tunnel_status,omitempty"`
	// Management connection status
	MgmtConnectionStatus string `json:"mgmt_connection_status,omitempty"`
	ControlConnectionStatus *StatusCount `json:"control_connection_status,omitempty"`
	PnicStatus *StatusCount `json:"pnic_status,omitempty"`
	NodeStatus *NodeStatus `json:"node_status,omitempty"`
	// Transport node display name
	NodeDisplayName string `json:"node_display_name,omitempty"`
}
