package nsxt

type ClusterNodeInfo struct {
	// Node clustering status
	Status string `json:"status,omitempty"`
	// Messaging client of all entities
	MsgClients []NodeMessagingClientInfo `json:"msg_clients,omitempty"`
	// Unique identifier of this node
	NodeUuid string `json:"node_uuid,omitempty"`
	// The display name of this node
	DisplayName string `json:"display_name,omitempty"`
	// Service endpoint of all entities
	Entities []NodeEntityInfo `json:"entities,omitempty"`
	// Certificate and thumbprint of all entities
	Certificates []NodeCertificateInfo `json:"certificates,omitempty"`
	// The fqdn of this node
	Fqdn string `json:"fqdn,omitempty"`
}
