package nsxt

type ManagementClusterRoleConfig struct {
	// Type of this role configuration
	Type_ string `json:"type,omitempty"`
	MgmtClusterListenAddr *ServiceEndpoint `json:"mgmt_cluster_listen_addr,omitempty"`
	MpaMsgClientInfo *MsgClientInfo `json:"mpa_msg_client_info,omitempty"`
	ApiListenAddr *ServiceEndpoint `json:"api_listen_addr,omitempty"`
	ApplianceConnectionInfo *ServiceEndpoint `json:"appliance_connection_info,omitempty"`
	MgmtPlaneListenAddr *ServiceEndpoint `json:"mgmt_plane_listen_addr,omitempty"`
}
