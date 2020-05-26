package nsxt

type EdgeClusterMemberTransportNode struct {
	// System generated index for cluster member
	MemberIndex int32 `json:"member_index"`
	// Identifier of the transport node backed by an Edge node
	TransportNodeId string `json:"transport_node_id"`
}
