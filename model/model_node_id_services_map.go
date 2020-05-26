package nsxt

type NodeIdServicesMap struct {
	// List of ServiceTypes.
	ServiceTypes []string `json:"service_types"`
	// NodeId
	NodeId string `json:"node_id"`
}
