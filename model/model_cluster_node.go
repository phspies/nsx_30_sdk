package nsxt

// This type contains attributes of a cluster node that are relevant to the Cluster Boot Manager.
type ClusterNode struct {
	// Current clustering status of the node
	Status string `json:"status,omitempty"`
	// Entities on the node
	Entities []ClusterNodeEntity `json:"entities"`
	// UUID of the node
	NodeUuid string `json:"node_uuid"`
}
