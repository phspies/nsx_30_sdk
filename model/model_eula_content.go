package nsxt

// End User License Agreement content
type EulaContent struct {
	Self *SelfResourceLink `json:"_self,omitempty"`
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink `json:"_links,omitempty"`
	// Schema for this resource
	Schema string `json:"_schema,omitempty"`
	// Content of End User License Agreement
	Content string `json:"content,omitempty"`
}
