package nsxt

// Indicate the status of End User License Agreement acceptance
type EulaAcceptance struct {
	Self *SelfResourceLink `json:"_self,omitempty"`
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink `json:"_links,omitempty"`
	// Schema for this resource
	Schema string `json:"_schema,omitempty"`
	// Acceptance status of End User License Agreement
	Acceptance bool `json:"acceptance"`
}
