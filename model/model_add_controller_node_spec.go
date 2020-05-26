package nsxt

type AddControllerNodeSpec struct {
	MpaMsgClientInfo *MsgClientInfo `json:"mpa_msg_client_info"`
	HostMsgClientInfo *MsgClientInfo `json:"host_msg_client_info"`
	ClusteringParams *ClusteringInfo `json:"clustering_params,omitempty"`
	// Only use this if an id for the node already exists with MP. If not specified, then the node_id will be set to a random id.
	NodeId string `json:"node_id,omitempty"`
	// Deprecated. Do not supply a value for this property.
	ControlPlaneServerCertificate string `json:"control_plane_server_certificate,omitempty"`
	// must be set to AddControllerNodeSpec
	Type_ string `json:"type"`
}
