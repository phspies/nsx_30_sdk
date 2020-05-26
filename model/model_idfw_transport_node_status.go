package nsxt

// ID and status of the Identity Firewall enabled Compute collection's transport node. 
type IdfwTransportNodeStatus struct {
	// Status of the IDFW transport node.
	TransportNodeStatus []IdfwTransportNodeCondition `json:"transport_node_status"`
	// TransportNode ID of the Identity Firewall enabled Compute collection's transport node. 
	TransportNodeId string `json:"transport_node_id"`
}
