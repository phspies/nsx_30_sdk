package nsxt

type VtepLabelRange struct {
	Self *SelfResourceLink `json:"_self,omitempty"`
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink `json:"_links,omitempty"`
	// Schema for this resource
	Schema string `json:"_schema,omitempty"`
	// Start value for virtual tunnel endpoint label range
	Start int64 `json:"start"`
	// End value for virtual tunnel endpoint label range
	End int64 `json:"end"`
}
