package nsxt

type NodeSummary struct {
	Self *SelfResourceLink `json:"_self,omitempty"`
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink `json:"_links,omitempty"`
	// Schema for this resource
	Schema string `json:"_schema,omitempty"`
	// Number of nodes of the type and at the component version
	NodeCount int32 `json:"node_count,omitempty"`
	// Node type
	Type_ string `json:"type,omitempty"`
	// Component version
	ComponentVersion string `json:"component_version,omitempty"`
}