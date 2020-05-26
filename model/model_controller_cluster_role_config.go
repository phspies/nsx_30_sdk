package nsxt

type ControllerClusterRoleConfig struct {
	// Type of this role configuration
	Type_ string `json:"type,omitempty"`
	MpaMsgClientInfo *MsgClientInfo `json:"mpa_msg_client_info,omitempty"`
	HostMsgClientInfo *MsgClientInfo `json:"host_msg_client_info,omitempty"`
	ControlPlaneListenAddr *ServiceEndpoint `json:"control_plane_listen_addr,omitempty"`
	ControlClusterListenAddr *ServiceEndpoint `json:"control_cluster_listen_addr,omitempty"`
}
