package nsxt

type NodeInfo struct {
	Self *SelfResourceLink `json:"_self,omitempty"`
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink `json:"_links,omitempty"`
	// Schema for this resource
	Schema string `json:"_schema,omitempty"`
	// Node type
	Type_ string `json:"type,omitempty"`
	// Name of the node
	DisplayName string `json:"display_name,omitempty"`
	// Identifier of the node
	Id string `json:"id,omitempty"`
	// Component version of the node
	ComponentVersion string `json:"component_version,omitempty"`
}
