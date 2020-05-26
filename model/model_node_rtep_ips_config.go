package nsxt

type NodeRtepIpsConfig struct {
	// System generated index for cluster member
	MemberIndex int32 `json:"member_index,omitempty"`
	// Remote tunnel endpoint ip address.
	RtepIps []string `json:"rtep_ips,omitempty"`
	// Identifier of the transport node backed by an Edge node
	TransportNodeId string `json:"transport_node_id,omitempty"`
}
