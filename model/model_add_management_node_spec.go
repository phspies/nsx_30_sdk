package nsxt

type AddManagementNodeSpec struct {
	MpaMsgClientInfo *MsgClientInfo `json:"mpa_msg_client_info,omitempty"`
	// must be set to AddManagementNodeSpec
	Type_ string `json:"type"`
	// The password to be used to authenticate with the remote node.
	Password string `json:"password,omitempty"`
	// The username to be used to authenticate with the remote node.
	UserName string `json:"user_name"`
	// The host address of the remote node to which to send this join request.
	RemoteAddress string `json:"remote_address"`
	// The certificate thumbprint of the remote node.
	CertThumbprint string `json:"cert_thumbprint,omitempty"`
}