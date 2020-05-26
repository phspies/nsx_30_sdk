package nsxt

// Status for IPSec VPN IKE session UP, DOWN, NEGOTIATING and fail reason if IKE session is down.
type IpSecVpnikeSessionStatus struct {
	// Transport Node identifier where session is present.
	TransportNodeId string `json:"transport_node_id,omitempty"`
	// Reason for failure.
	FailReason string `json:"fail_reason,omitempty"`
	// IKE session service status UP, DOWN and NEGOTIATING.
	IkeSessionState string `json:"ike_session_state,omitempty"`
}
