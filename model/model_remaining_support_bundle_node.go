package nsxt

type RemainingSupportBundleNode struct {
	// Status of node
	Status string `json:"status,omitempty"`
	// Display name of node
	NodeDisplayName string `json:"node_display_name,omitempty"`
	// UUID of node
	NodeId string `json:"node_id,omitempty"`
}
