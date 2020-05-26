package nsxt

type LogicalPortMirrorSource struct {
	// Resource types of mirror source
	ResourceType string `json:"resource_type"`
	// Source logical port identifier list
	PortIds []string `json:"port_ids"`
}
