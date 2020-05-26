package nsxt

// A link to a related resource
type ResourceLink struct {
	// Optional action
	Action string `json:"action,omitempty"`
	// Link to resource
	Href string `json:"href,omitempty"`
	// Custom relation type (follows RFC 5988 where appropriate definitions exist)
	Rel string `json:"rel,omitempty"`
}