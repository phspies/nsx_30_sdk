package nsxt

type LogicalPortMirrorDestination struct {
	// Resource types of mirror destination
	ResourceType string `json:"resource_type"`
	// Destination logical port identifier list.
	PortIds []string `json:"port_ids"`
}
